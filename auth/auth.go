package auth

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("test")

func CreateJwtToken(res http.ResponseWriter, req *http.Request) {
	username := req.Header.Get("User")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Minute * 1).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(res, tokenString)
}

func VerifyToken(authToken string) bool {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		fmt.Println(err)
		return false
	}

	if !token.Valid {
		fmt.Println("Invalid Token")
		return false
	}

	return true
}
