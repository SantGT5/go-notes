package router

import (
	"os"

	"github.com/SantGT5/quintosgo/api/handlers"
	"github.com/SantGT5/quintosgo/api/handlers/validate"
	"github.com/SantGT5/quintosgo/api/middleware"
)

// InitRoutes initializes all routes for the application
func Routes() {
	baseURLV1 := os.Getenv("BASE_URL_V1")

	v1 := router.Group(baseURLV1)
	{
		authRouter := v1.Group("/auth")
		{
			authRouter.POST("/login", handlers.Login)
		}

		user := v1.Group("/user")
		{
			user.POST("", handlers.PostUser)
			user.GET("/me", middleware.RequiredAuth, validate.Validate)
		}

	}
}
