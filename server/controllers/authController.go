package controllers

import (
	"log"

	"github.com/formatkaka/youtube-clone/db"
	"github.com/formatkaka/youtube-clone/helpers"
	"github.com/formatkaka/youtube-clone/models"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var req models.SignUpRequest

	if err := ctx.BindJSON(&req); err != nil {
		log.Fatal("bind issue signup request obj")
		ctx.JSON(500, gin.H{
			"error": "bind issue signup request obj",
		})
		return
	}

	hash, err := helpers.GenerateHash(req.Password)

	if err != nil {
		log.Println(err)
		ctx.JSON(200, gin.H{
			"error": err,
		})
		return
	}

	log.Println("hash : ", hash)

	client := db.GetClient()
	var id int64
	query := "INSERT INTO USERS (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err = client.QueryRow(query, req.Name, req.Email, hash).Scan(&id)

	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"err": err,
		})
		return
	}

	token, refreshToken, err := helpers.GenerateToken(req.Email, req.Name, id)

	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"err": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"token":        token,
		"refreshToken": refreshToken,
	})

}

func Login(ctx *gin.Context) {

}
