package main

import "github.com/labstack/echo/v4"

func SetupRoutes(app *echo.Echo, a *apiConfig) {
	group := app.Group("/v1")
	group.GET("/readiness", HandleReadiness)
	group.GET("/liveness", HandleLiveness)
	group.POST("/users", func(c echo.Context) error {
		return HandleCreateUser(c, a)
	})
	group.GET("/users", func(c echo.Context) error {
		return HandleGetUserByAPIKey(c, a)
	})
	group.POST("/feeds", func(c echo.Context) error {
		return HandleCreateFeed(c, a)
	})
	group.GET("/feeds", func(c echo.Context) error {
		return HandleGetFeeds(c, a)
	})
	group.GET("/feed_follows", func(c echo.Context) error {
		return HandleGetFeedFollows(c, a)
	})
	group.POST("/feed_follows", func(c echo.Context) error {
		return HandleCreateFeedFollow(c, a)
	})
	group.DELETE("/feed_follows/:feedFollowID", func(c echo.Context) error {
		return HandleDeleteFeedFollow(c, a)
	})
	group.GET("/posts", func(c echo.Context) error {
		return HandlerGetPostsForUser(c, a)
	})
}
