package main

import (
	"fmt"

	"go.xstore.local/go-hit-counter/config"
	"go.xstore.local/go-hit-counter/database"
	"go.xstore.local/go-hit-counter/routes"
)

func main() {
	if err := config.Load("config/config.yaml"); err != nil {
		fmt.Println("Failed to load configuration")
		return
	}

	db, err := database.InitDB()
	if err != nil {
		fmt.Println("err open databases")
		return
	}
	defer db.Close()

	router := routes.InitRouter()
	router.Run(config.Get().Addr)
}
