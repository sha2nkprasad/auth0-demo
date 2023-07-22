package api

import (
	"auth0-example/auth"
	"auth0-example/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func handleRoutes() gin.HandlerFunc {
	return func(context *gin.Context) {
		userController := controllers.NewUserController()
		userController.Profile(context)
	}
}

func Router() *gin.Engine {
	// Setting up Gin
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(auth.GetUser)

	r.POST("/user", handleRoutes())
	return r
}
