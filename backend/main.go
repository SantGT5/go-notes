package main

import (
	"fmt"

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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run()

	if err != nil {
		fmt.Println("Error starting server:", err.Error())
	}
}
