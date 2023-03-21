package mw

import (
	"gotodo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Bouncer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, cookieErr := ctx.Cookie("auth")
		if cookieErr != nil || cookie == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized."})
			return
		}

		claims, parsingErr := utils.ValidateJWT(cookie)
		if parsingErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed. Malformed token."})
			return
		}
		ctx.Set("claims", claims)

		ctx.Next()
	}
}
