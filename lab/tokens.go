package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	Email string
	Uid   string
	jwt.RegisteredClaims
}

func main() {
	authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": "aaryakdshah@gmail.com",
		"uid":   "1234-5678-abcd-ef90",
	})
	authTokenStr, signingErr := authToken.SignedString([]byte("lab_secret_string"))
	if signingErr != nil {
		fmt.Printf("token=%#v\nsecret=%s\n", authToken, "lab_secret_string")
		fmt.Printf("s_error=%#v\n", signingErr)
	}
	fmt.Printf("signed=%s\n", authTokenStr)

	claims := jwt.MapClaims{}
	token, parsingErr := jwt.ParseWithClaims(authTokenStr, claims, func(token *jwt.Token) (interface{}, error) { return []byte("lab_secret_string"), nil })
	if parsingErr != nil {
		fmt.Printf("p_error=%#v\n", parsingErr)
		return
	}
	fmt.Printf("token=%#v\n", token)
	fmt.Printf("email=%s\n", claims["email"])
	fmt.Printf("uid=%s\n", claims["uid"])
}
