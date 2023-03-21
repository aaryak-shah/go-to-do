package routes

import (
	"fmt"
	"gotodo/db"
	"gotodo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ViewAll(ctx *gin.Context) {
	fmt.Printf("ctxkeys=%#v\n", ctx.Keys)
	uid := ctx.Value("claims").(jwt.MapClaims)["uid"].(string)
	user := models.User{}
	result := db.Instance.First(&user, "uid = ?", uid)
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't fetch data. Try again later."})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user.ToDos})
}

func ViewId(ctx *gin.Context) {}
