package main

import (
	"awesomeProject/db"
	"awesomeProject/server"
	"fmt"
)

func main() {

	err := db.Connect()
	if err != nil {
		fmt.Println("Error connecting to database")
	}
	server.Init()
}
