package routes

import (
	"github.com/formatkaka/youtube-clone/controllers"
	"github.com/formatkaka/youtube-clone/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	routes.Use(middleware.Authenticate)
	routes.GET("/user", controllers.GetUserInfo)
}
