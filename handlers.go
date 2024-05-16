package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func respondWithError(c echo.Context, code int, err error) error {
	if code > 499 {
		log.Println("Responding with 5xx error: ", err)
	}
	return c.JSON(code, map[string]interface{}{
		"error": err.Error(),
	})
}

func HandleReadiness(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func HandleLiveness(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func HandleCreateUser(c echo.Context, a *apiConfig) error {
	return a.CreateUser(c)
}

func HandleGetUserByAPIKey(c echo.Context, a *apiConfig) error {
	return a.GetUserByAPIKey(c)
}

func HandleCreateFeed(c echo.Context, a *apiConfig) error {
	user, err := a.ValidateAPIKey(c)
	if err != nil {
		if err.Error() == "API key is missing" {
			return respondWithError(c, http.StatusUnauthorized, err)
		}
		if err == sql.ErrNoRows {
			return respondWithError(c, http.StatusUnauthorized, fmt.Errorf("invalid API key"))
		}
		return respondWithError(c, http.StatusInternalServerError, err)
	}
	return a.CreateFeed(c, user)
}

func HandleGetFeeds(c echo.Context, a *apiConfig) error {
	return a.GetFeeds(c)
}

func HandleCreateFeedFollow(c echo.Context, a *apiConfig) error {
	user, err := a.ValidateAPIKey(c)
	if err != nil {
		if err.Error() == "API key is missing" {
			return respondWithError(c, http.StatusUnauthorized, err)
		}
		if err == sql.ErrNoRows {
			return respondWithError(c, http.StatusUnauthorized, fmt.Errorf("invalid API key"))
		}
		return respondWithError(c, http.StatusInternalServerError, err)
	}
	return a.CreateFeedFollow(c, user)
}

func HandleDeleteFeedFollow(c echo.Context, a *apiConfig) error {
	user, err := a.ValidateAPIKey(c)
	if err != nil {
		if err.Error() == "API key is missing" {
			return respondWithError(c, http.StatusUnauthorized, err)
		}
		if err == sql.ErrNoRows {
			return respondWithError(c, http.StatusUnauthorized, fmt.Errorf("invalid API key"))
		}
		return respondWithError(c, http.StatusInternalServerError, err)
	}
	return a.DeleteFeedFollow(c, user)
}

func HandleGetFeedFollows(c echo.Context, a *apiConfig) error {
	user, err := a.ValidateAPIKey(c)
	if err != nil {
		if err.Error() == "API key is missing" {
			return respondWithError(c, http.StatusUnauthorized, err)
		}
		if err == sql.ErrNoRows {
			return respondWithError(c, http.StatusUnauthorized, fmt.Errorf("invalid API key"))
		}
		return respondWithError(c, http.StatusInternalServerError, err)
	}
	return a.GetFeedFollows(c, user)
}

func HandlerGetPostsForUser(c echo.Context, a *apiConfig) error {
	user, err := a.ValidateAPIKey(c)
	if err != nil {
		if err.Error() == "API key is missing" {
			return respondWithError(c, http.StatusUnauthorized, err)
		}
		if err == sql.ErrNoRows {
			return respondWithError(c, http.StatusUnauthorized, fmt.Errorf("invalid API key"))
		}
		return respondWithError(c, http.StatusInternalServerError, err)
	}
	return a.GetPostsForUser(c, user)
}
