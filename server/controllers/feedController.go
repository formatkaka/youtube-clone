package controllers

import (
	"log"

	"github.com/formatkaka/youtube-clone/db"
	"github.com/formatkaka/youtube-clone/models"
	"github.com/gin-gonic/gin"
)

func GetFeed(ctx *gin.Context) {
	client := db.GetClient()

	query := `SELECT id, title, video_url, poster, duration, views, user_id, created_at  FROM videos LIMIT 15`
	rows, err := client.Query(query)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var feed []models.Video

	for rows.Next() {
		var video models.Video
		if err := rows.Scan(&video.ID, &video.Title, &video.Video_url, &video.Poster, &video.Duration, &video.Views, &video.User_id, &video.Created_at); err != nil {
			panic(err)
		}
		feed = append(feed, video)
	}

	ctx.JSON(200, gin.H{
		"response": feed,
	})
}
