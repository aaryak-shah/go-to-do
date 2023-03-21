package main

import (
	"gotodo/db"
	"gotodo/initializers"
	"gotodo/mw"
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

	corsPolicy := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "User-Agent"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsPolicy))

	api := router.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })

	auth := v1.Group("/auth")
	auth.POST("/register", routes.Register)
	auth.POST("/login", routes.Login)
	auth.GET("/logout", routes.Logout)
	auth.GET("/check", routes.CheckAuth)

	create := v1.Group("/create")
	create.Use(mw.Bouncer())
	create.POST("/", routes.Create)

	view := v1.Group("/view")
	view.Use(mw.Bouncer())
	view.GET("/", routes.ViewAll)
	view.GET("/:id", routes.ViewId)

	edit := v1.Group("/edit")
	edit.Use(mw.Bouncer())
	edit.PATCH("/:id", routes.Edit)

	// Start Server
	router.Run("localhost:9000")
}
