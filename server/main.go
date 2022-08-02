package main

import "github.com/gin-gonic/gin"

func main() {
	port := "8080"

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"response": "pong"})
	})

	router.Run(":" + port)
}
