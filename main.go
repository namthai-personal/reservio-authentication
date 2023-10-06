package main

import (
	"fmt"
	"reservio-be/infrastructure/database"
	"reservio-be/routes"
)

const (
	PORT = ":3000"
)

func main() {
	fmt.Println("Init Application")

	router := routes.SetupRouter()
	database.OpenDbConnection()
	router.Run(PORT)
}
