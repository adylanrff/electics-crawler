package model

import "time"

type Tweet struct {
	TweetID       int64      `json:"tweet_id"`
	Timestamp     time.Time  `json:"timestamp"`
	RetweetCount  int        `json:"retweet_count"`
	FavoriteCount int        `json:"favorite_count"`
	Text          string     `json:"text"`
	Coordinates   [2]float64 `json:"coordinates"`
	Username      string     `json:"username"`
}
