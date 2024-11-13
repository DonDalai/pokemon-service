package main

import (
	"pokemon-service/database"
	"pokemon-service/routes"
)

func main() {
	database.ConnectDatabase()
	database.Migrate()

	router := routes.SetupRouter()
	router.Run(":8080")
}
