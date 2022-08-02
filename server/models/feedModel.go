package models

import "time"

type Video struct {
	ID         int       `json: "id"`
	Title      string    `json: "title"`
	Video_url  string    `json: "video_url"`
	Poster     string    `json: "poster"`
	User_id    int       `json: "user_id"`
	Duration   int       `json: "duration"`
	Views      int       `json: "views"`
	Created_at time.Time `json: "created_at"`
}

type Feed struct {
	feed []Video
}
