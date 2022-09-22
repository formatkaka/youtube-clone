package controllers

import "github.com/gin-gonic/gin"

func GetUserInfo(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"name":  ctx.GetString("name"),
		"email": ctx.GetString("email"),
	})
}
