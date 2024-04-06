package main

import (
	"github.com/SantGT5/quintosgo/database"
	"github.com/SantGT5/quintosgo/router"
)

func main() {
	database.Init()

	router.Init()
}
