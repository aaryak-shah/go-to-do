package main

import (
	"gotodo/db"
	"gotodo/initializers"
	"gotodo/routes"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	// Load Environment
	initializers.LoadEnv()

	// Connect to Postgres Instance
	db.AttachInstance()
}

func main() {
	// Set up Routing
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:9000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "User-Agent"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })

	auth := v1.Group("/auth")
	auth.POST("/register", routes.Register)
	auth.POST("/login", routes.Login)
	auth.GET("/logout", routes.Logout)
	auth.GET("/check", routes.CheckAuth)

	v1.POST("/create", routes.Create)

	view := v1.Group("/view")
	view.GET("/", routes.ViewAll)
	view.GET("/:id", routes.ViewId)

	v1.PATCH("/edit/:id", routes.Edit)

	// Start Server
	router.Run("localhost:9000")
}
