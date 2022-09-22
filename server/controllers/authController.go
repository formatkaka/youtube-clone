package controllers

import (
	"fmt"
	"log"

	"github.com/formatkaka/youtube-clone/db"
	"github.com/formatkaka/youtube-clone/helpers"
	"github.com/formatkaka/youtube-clone/models"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var req models.SignUpRequest

	if err := ctx.BindJSON(&req); err != nil {
		log.Println("bind issue signup request obj : ", err)
		ctx.JSON(500, gin.H{
			"error": fmt.Sprintf("bind issue signup request obj :  %s", err.Error()),
		})
		return
	}

	hash, err := helpers.GenerateHash(req.Password)

	if err != nil {
		log.Println("generate hash error in signup func ", err)
		ctx.JSON(200, gin.H{
			"error": fmt.Sprintf("generate hash error in signup func %s ", err),
		})
		return
	}

	client := db.GetClient()
	var id int64
	query := "INSERT INTO USERS (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err = client.QueryRow(query, req.Name, req.Email, hash).Scan(&id)

	if err != nil {
		// how to return error that email already exists in db ??
		// should I check if email exists and then put or try to analyse the error message
		log.Println("unable to insert ", err)
		ctx.JSON(500, gin.H{
			"error": fmt.Sprintf("unable to insert %s ", err),
		})
		return
	}

	token, refreshToken, err := helpers.GenerateToken(req.Email, req.Name, id)

	if err != nil {
		log.Println("signup request : ", err)
		ctx.JSON(500, gin.H{
			"error":  fmt.Sprintf("signup request : %s", err),
			"status": helpers.TOKEN_ERROR,
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
		log.Println("bind issue login request obj : ", err)
		ctx.JSON(500, gin.H{
			"error": fmt.Sprintf("bind issue login request obj :  %s", err.Error()),
		})
		return
	}

	fmt.Println("req : ", req)

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
			log.Println("ERROR: login rows Scan : ", err)
		}
	} else {
		ctx.JSON(500, gin.H{
			"error":  "Email does not exist",
			"status": helpers.EMAIL_NOT_EXIST,
		})
		return
	}

	userAuthenticated := helpers.CompareHash(req.Password, hash)
	token, refreshToken, err := helpers.GenerateToken(req.Email, name, id)

	if err != nil {
		log.Println("login func : ", err)
		ctx.JSON(500, gin.H{
			"error":  fmt.Sprintf("login function token : %s", err.Error()),
			"status": helpers.TOKEN_ERROR,
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

	ctx.JSON(401, gin.H{
		"error":     "Invalid email or password",
		"errorcode": helpers.INVALID_LOGIN_CREDS,
	})

}
