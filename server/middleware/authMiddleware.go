package middleware

import (
	"net/http"

	"github.com/formatkaka/youtube-clone/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {

	clientToken := ctx.Request.Header.Get("token")

	if clientToken == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Auth headers not provided",
		})
		ctx.Abort()
		return
	}

	claims, err := helpers.ValidateToken(clientToken)

	if err != "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		ctx.Abort()
		return
	}

	ctx.Set("email", claims.Email)
	ctx.Set("name", claims.Name)
	ctx.Set("uid", claims.Uid)

	ctx.Next()

}
