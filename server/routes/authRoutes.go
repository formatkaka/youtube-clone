package routes

import (
	"github.com/formatkaka/youtube-clone/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(routes *gin.Engine) {
	routes.POST("signup", controllers.SignUp)
	routes.POST("login", controllers.Login)
}
