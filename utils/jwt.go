package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func SignJWT(claims jwt.MapClaims) (string, error) {
	authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	authTokenStr, signingErr := authToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if signingErr != nil {
		return "", signingErr
	}
	return authTokenStr, nil
}

func ValidateJWT(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, parsingErr := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) { return []byte(os.Getenv("JWT_SECRET")), nil })
	if parsingErr != nil {
		return claims, parsingErr
	}
	return claims, nil
}
