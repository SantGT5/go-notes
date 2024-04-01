package main

import (
	"github.com/SantGT5/notes/initializers"
	"github.com/SantGT5/notes/router"
)

func int() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()

	router.Initializer()
}

func main() {
	int()
}
