package main

import (
	"github.com/SantGT5/quintosgo/internal/database"
	"github.com/SantGT5/quintosgo/internal/router"
)

func main() {
	database.Init()
	router.Init()
}
