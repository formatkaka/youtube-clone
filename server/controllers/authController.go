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
	var req models.LoginRequest

	if err := ctx.BindJSON(&req); err != nil {
		log.Fatal("bind issue signup request obj")
		ctx.JSON(500, gin.H{
			"error": "bind issue signup request obj",
		})
		return
	}

	client := db.GetClient()
	var name, hash string
	var id int64
	query := "SELECT id, name, password FROM USERS WHERE email = $1"

	rows, err := client.Query(query, req.Email)

	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"error": err,
		})
	}

	if rows.Next() {
		if err := rows.Scan(&id, &name, &hash); err != nil {
			panic(err)
		}
	} else {
		ctx.JSON(500, gin.H{
			"error": "Email does not exist",
		})
		return
	}

	userAuthenticated := helpers.CompareHash(req.Password, hash)
	token, refreshToken, err := helpers.GenerateToken(req.Email, name, id)

	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"err": err,
		})
		return
	}

	if userAuthenticated {
		ctx.JSON(200, gin.H{
			"token":        token,
			"refreshToken": refreshToken,
			"name":         name,
			"email":        req.Email,
		})
		return
	}

	ctx.JSON(500, gin.H{
		"error": "Invalid email or password",
	})

}
