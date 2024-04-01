package router

import (
	"os"

	"github.com/SantGT5/notes/controllers"
	"github.com/SantGT5/notes/middleware"
	"github.com/gin-gonic/gin"
)

// InitRoutes initializes all routes for the application
func InitRoutes(router *gin.Engine) {
	baseURLV1 := os.Getenv("BASE_URL_V1")

	v1 := router.Group(baseURLV1)
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
		}

		user := v1.Group("/user")
		{
			user.POST("/signup", controllers.Signup)
			user.GET("/me", middleware.RequiredAuth, controllers.Validate)
		}

	}
}
