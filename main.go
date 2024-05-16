package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hackertron/blog-agg/internal/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("blog go brrr")
	godotenv.Load(".env")
	PORT := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not set")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	apiConfig := NewApiConfig(database.New(conn))

	go startScraping(apiConfig.DB, 10, time.Minute)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		// send json response
		return c.JSON(http.StatusOK, "hello")
	})
	SetupRoutes(e, apiConfig)
	e.Logger.Fatal(e.Start("localhost:" + PORT))
}
