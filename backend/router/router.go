package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Initializer() {
	var router *gin.Engine = gin.Default()

	// Cors Config
	Cors(router)

	// init Routers
	InitRoutes(router)

	port := os.Getenv("BACKEND_PORT")

	err := router.Run(":" + port)

	if err != nil {
		fmt.Println("Error starting server:", err.Error())
	}
}
