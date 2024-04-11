package router

import (
	"os"

	"github.com/SantGT5/quintosgo/api/handlers/auth"
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
			authRouter.POST("/login", auth.Login)
			authRouter.POST("/signup", auth.Signup)
		}

		user := v1.Group("/user")
		{
			user.GET("/me", middleware.RequiredAuth, validate.Validate)
		}

	}
}
