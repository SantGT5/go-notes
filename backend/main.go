package main

import (
	"fmt"
	"os"

	"github.com/SantGT5/notes/controllers"
	"github.com/SantGT5/notes/initializers"
	"github.com/gin-gonic/gin"
)

func int() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	int()

	r := gin.Default()
	r.POST("/signup", controllers.Signup)

	addr := ":" + os.Getenv("BACKEND_PORT")

	err := r.Run(addr)

	if err != nil {
		fmt.Println("Error starting server:", err.Error())
	}
}
