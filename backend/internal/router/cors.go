package router

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
)

func Cors() {
	allowOrigins := os.Getenv("ALLOW_ORIGINS")
	origins := strings.Split(allowOrigins, ",")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // Maximum age for cache (in seconds)
	}))
}
