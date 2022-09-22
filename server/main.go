package main

import (
	"github.com/formatkaka/youtube-clone/db"
	"github.com/formatkaka/youtube-clone/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())
	db.Init()

	routes.FeedRoutes(router)
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"response": "pong"})
	})

	router.Run(":" + port)
}
