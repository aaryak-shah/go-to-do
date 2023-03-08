package main

import (
	"gotodo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/register", routes.Register)
	auth.POST("/login", routes.Login)
	auth.GET("/logout", routes.Logout)

	v1.POST("/create", routes.Create)

	view := v1.Group("/view")
	view.GET("/", routes.ViewAll)
	view.GET("/:id", routes.ViewId)

	v1.PATCH("/edit/:id", routes.Edit)

	router.Run("localhost:9000")
}
