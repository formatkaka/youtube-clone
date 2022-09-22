package models

import "time"

type FeedObject struct {
	Title      string    `json: "title"`
	Video_url  string    `json: "video_url"`
	Poster     string    `json: "poster"`
	Name       string    `json: "name"`
	Duration   int       `json: "duration"`
	Views      int       `json: "views"`
	Created_at time.Time `json: "created_at"`
}

type Feed struct {
	feed []FeedObject
}
