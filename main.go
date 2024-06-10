package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/shashank855/server/auth"
)

func getResposnse(res http.ResponseWriter, req *http.Request) {
	authToken := req.Header.Get("Authentication")
	if authToken == "" {
		io.WriteString(res, "Authentication misssing\n")
	} else {
		if auth.VerifyToken(authToken) {
			io.WriteString(res, "hello from server\n")
		} else {
			io.WriteString(res, "Token Invalid\n")
		}
	}
}

func main() {
	http.HandleFunc("/", getResposnse)
	http.HandleFunc("/createToken", auth.CreateJwtToken)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println(err)
	}
}
