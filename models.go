package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/hackertron/blog-agg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		APIKey:    user.ApiKey,
	}
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

// return feeds as database feeds
// func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
// 	var databaseFeeds []Feed
// 	for _, feed := range feeds {
// 		databaseFeeds = append(databaseFeeds, databaseFeedToFeed(feed))
// 	}
// 	return databaseFeeds
// }

// func databaseFeedFollowToFeedFollow(feedFollow database.FeedsFollow) FeedFollow {
// 	return FeedFollow{
// 		ID:        feedFollow.ID,
// 		CreatedAt: feedFollow.CreatedAt,
// 		UpdatedAt: feedFollow.UpdatedAt,
// 		UserID:    feedFollow.UserID,
// 		FeedID:    feedFollow.FeedID,
// 	}
// }

// // return feeds as database feeds
// func databaseFeedFollowsToFeedFollows(feedFollows []database.FeedsFollow) []FeedFollow {
// 	var databaseFeedFollows []FeedFollow
// 	for _, feedFollow := range feedFollows {
// 		databaseFeedFollows = append(databaseFeedFollows, databaseFeedFollowToFeedFollow(feedFollow))
// 	}
// 	return databaseFeedFollows
// }
