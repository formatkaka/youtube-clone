package helpers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email string
	Name  string
	Uid   int64
	jwt.StandardClaims
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateToken(email string, name string, uid int64) (string, string, error) {
	claims := &SignedDetails{
		Email: email,
		Name:  name,
		Uid:   uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Println(err)
		return "", "", fmt.Errorf("token generation error in GenerateToken : %w", err)
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Println(err)
		return "", "", fmt.Errorf("refresh token generation error in GenerateToken : %w", err)
	}

	return token, refreshToken, nil

}
