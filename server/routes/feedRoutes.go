package routes

import (
	"github.com/formatkaka/youtube-clone/controllers"
	"github.com/gin-gonic/gin"
)

func FeedRoutes(routes *gin.Engine) {
	routes.GET("/feed", func(ctx *gin.Context) {
		controllers.GetFeed(ctx)
	})
}
