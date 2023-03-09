package routes

import (
	"fmt"
	"gotodo/db"
	"gotodo/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	ctx.Bind(&body)

	uid := uuid.NewString()

	passwordHash, hashErr := bcrypt.GenerateFromPassword([]byte(body.Password), 0)
	if hashErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to register. Try again later."})
		fmt.Println(hashErr)
		return
	}

	user := models.User{}
	checkExisting := db.Instance.Where("email = ?", body.Email).First(&user)
	fmt.Printf("CHECKEXISTING %#v", checkExisting)
	if checkExisting.RowsAffected > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("User with email %s alerady exists.", user.Email)})
		return
	}

	user = models.User{UID: uid, Email: body.Email, Password: string(passwordHash)}
	result := db.Instance.Create(&user)
	if result.Error != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	// authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"email": user.Email,
	// 	"uid":   user.UID,
	// })
	// authTokenStr, signingErr := authToken.SignedString(os.Getenv("JWT_SECRET"))
	// if signingErr != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to process login. Try again later."})
	// }
	// ctx.SetCookie("auth", authTokenStr, 3600*24*30, "/", os.Getenv("TARGET_HOST"), false, false)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Registered successfully."})
}

func Login(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	ctx.Bind(&body)

	user := models.User{Email: body.Email}
	findUser := db.Instance.Where("email = ?", body.Email).First(&user)
	if findUser.RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No user with this email."})
		return
	} else {
		passwordMatchErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
		if passwordMatchErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password."})
			return
		}
	}

	authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"uid":   user.UID,
	})
	authTokenStr, signingErr := authToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if signingErr != nil {
		fmt.Printf("token=%#v\nsecret=%s\n", authToken, os.Getenv("JWT_SECRET"))
		fmt.Printf("error=%#v\n", signingErr)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to login. Try again later."})
		return
	}
	ctx.SetCookie("auth", authTokenStr, 3600*24*30, "/", os.Getenv("TARGET_HOST"), false, false)
	ctx.JSON(http.StatusCreated, gin.H{
		"email": user.Email,
		"uid":   user.UID,
	})
}

func Logout(ctx *gin.Context) {
	ctx.SetCookie("auth", "", -1, "/", os.Getenv("TARGET_HOST"), false, false)
}

func CheckAuth(ctx *gin.Context) {
	type JWTClaims struct {
		Email string
		Uid   string
		jwt.RegisteredClaims
	}

	cookie, err := ctx.Cookie("auth")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cookies not sent. Try again with cookies enabled."})
		return
	}

	token, err := jwt.ParseWithClaims(cookie, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) { return []byte(os.Getenv("JWT_SECRET")), nil })
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Try again later."})
		return
	}

}
