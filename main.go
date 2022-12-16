package main

import (
	"FifaDataProvider/database"
	"FifaDataProvider/routes"
	"fmt"
)

func main() {
	fmt.Println("Initiate Fifa Data Provider ...")

	database.Connect()
	routes.HandleRequest()
}
