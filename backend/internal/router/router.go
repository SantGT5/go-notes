package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() error {
	router = gin.Default()

	// Cors Config
	Cors()

	// init Routers
	Routes()

	port := os.Getenv("BACKEND_PORT")

	err := router.Run(":" + port)

	if err != nil {
		fmt.Println("Error starting server:", err.Error())
		return err
	}

	return nil
}
