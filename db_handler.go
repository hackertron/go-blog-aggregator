package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hackertron/blog-agg/internal/database"
	"github.com/labstack/echo/v4"
)

type apiConfig struct {
	DB *database.Queries
}

func NewApiConfig(db *database.Queries) *apiConfig {
	return &apiConfig{
		DB: db,
	}
}

func (a *apiConfig) ValidateAPIKey(c echo.Context) (database.User, error) {
	apiKey := c.Request().Header.Get("X-API-Key")
	if apiKey == "" {
		return database.User{}, fmt.Errorf("API key is missing")
	}

	user, err := a.DB.GetUserByAPIKey(c.Request().Context(), apiKey)
	if err != nil {
		return database.User{}, err
	}

	return user, nil
}

func (a *apiConfig) CreateUser(c echo.Context) error {
	type parameters struct {
		Name string `json:"name"`
	}
	// get json from echo context
	p := new(parameters)
	if err := c.Bind(p); err != nil {
		return respondWithError(c, http.StatusBadRequest, err)
	}

	log.Println("creating user with name: ", p.Name)

	// create user
	user, err := a.DB.CreateUser(c.Request().Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      p.Name,
	})
	if err != nil {
		return respondWithError(c, http.StatusInternalServerError, err)
	}
	// send response
	return c.JSON(http.StatusOK, databaseUserToUser(user))
}

func (a *apiConfig) CreateFeed(c echo.Context, user database.User) error {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	// get json from echo context
	p := new(parameters)
	if err := c.Bind(p); err != nil {
		return respondWithError(c, http.StatusBadRequest, err)
	}
	// create feed
	feed, err := a.DB.CreateFeed(c.Request().Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      p.Name,
		Url:       p.URL,
		UserID:    user.ID,
	})
	if err != nil {
		return respondWithError(c, http.StatusInternalServerError, err)
	}
	// create feed follow
	feed_follow, err := a.DB.CreateFeedFollow(c.Request().Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return respondWithError(c, http.StatusInternalServerError, err)
	}
	// send response
	resp := map[string]interface{}{
		"feed":        databaseFeedToFeed(feed),
		"feed_follow": feed_follow,
	}
	return c.JSON(http.StatusOK, resp)
}

func (a *apiConfig) GetUserByAPIKey(c echo.Context) error {
	type parameters struct {
		APIKey string `json:"api_key"`
	}
	// get api key from authorization headers
	apiKey := c.Request().Header.Get("X-API-Key")
	if apiKey == "" {
		return respondWithError(c, http.StatusBadRequest, fmt.Errorf("missing api key"))
	}
	// get json from echo context
	p := new(parameters)
	if err := c.Bind(p); err != nil {
		return respondWithError(c, http.StatusBadRequest, err)
	}
	// get user
	user, err := a.DB.GetUserByAPIKey(c.Request().Context(), apiKey)
	if err != nil {
		return respondWithError(c, http.StatusInternalServerError, err)
	}
	// send response
	return c.JSON(http.StatusOK, databaseUserToUser(user))
}

func (a *apiConfig) GetFeeds(c echo.Context) error {
	feeds, err := a.DB.GetFeeds(c.Request().Context())
	if err != nil {
		return respondWithError(c, http.StatusInternalServerError, err)
	}
	// send response
	return c.JSON(http.StatusOK, feeds)
}

func (a *apiConfig) CreateFeedFollow(c echo.Context, user database.User) error {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	// get json from echo context
	p := new(parameters)
	if err := c.Bind(p); err != nil {
		return respondWithError(c, http.StatusBadRequest, err)
	}
	log.Println("p : ", p)
	log.Println("feedID : ", p.FeedID)
	log.Println("userID : ", user.ID)
	// create feed follow
	feedFollow, err := a.DB.CreateFeedFollow(c.Request().Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    p.FeedID,
	})
	if err != nil {
		return respondWithError(c, http.StatusInternalServerError, err)
	}
	// send response
	// TODO : complete DatabaseFeedFollowToFeedFollow
	return c.JSON(http.StatusOK, feedFollow)
}

func (a *apiConfig) DeleteFeedFollow(c echo.Context, user database.User) error {
	feedFollowID := c.Param("feedFollowID")
	if feedFollowID == "" {
		return respondWithError(c, http.StatusBadRequest, fmt.Errorf("missing feed follow ID"))
	}
	// string to UUID
	feedFollowUUID, err := uuid.Parse(feedFollowID)
	if err != nil {
		return respondWithError(c, http.StatusBadRequest, err)
	}
	// delete feed follow
	Del_err := a.DB.DeleteFeedFollow(c.Request().Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feedFollowUUID,
	})
	if Del_err != nil {
		return respondWithError(c, http.StatusInternalServerError, err)
	}
	// send response
	return c.NoContent(http.StatusNoContent)
}

func (a *apiConfig) GetFeedFollows(c echo.Context, user database.User) error {
	feedFollows, err := a.DB.GetFeedFollows(c.Request().Context(), user.ID)
	if err != nil {
		return respondWithError(c, http.StatusInternalServerError, err)
	}
	// send response
	return c.JSON(http.StatusOK, feedFollows)
}

func (a *apiConfig) GetPostsForUser(c echo.Context, user database.User) error {
	posts, err := a.DB.GetPostsForUser(c.Request().Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		return respondWithError(c, http.StatusBadRequest, err)
	}
	// send response
	return c.JSON(http.StatusOK, posts)
}
