package controllers

import (
	"log"
	"strconv"

	"github.com/formatkaka/youtube-clone/db"
	"github.com/formatkaka/youtube-clone/models"
	"github.com/gin-gonic/gin"
)

func GetFeed(ctx *gin.Context) {
	page, ok := ctx.GetQuery("page")
	if !ok {
		ctx.JSON(500, gin.H{
			"error": "Some error occured with page param",
		})
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	limit, ok := ctx.GetQuery("limit")
	if !ok {
		ctx.JSON(500, gin.H{
			"error": "Some error occured with limit param",
		})
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	client := db.GetClient()

	query := `SELECT title, video_url, poster, name, duration, views, videos.created_at  
				FROM videos 
				JOIN users ON users.id = videos.user_id
				LIMIT $1 OFFSET $2`
	rows, err := client.Query(query, limitInt, pageInt*limitInt)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var feed []models.FeedObject

	for rows.Next() {
		var video models.FeedObject
		if err := rows.Scan(&video.Title, &video.Video_url, &video.Poster, &video.Name, &video.Duration, &video.Views, &video.Created_at); err != nil {
			panic(err)
		}
		feed = append(feed, video)
	}

	ctx.JSON(200, gin.H{
		"response": feed,
	})
}
