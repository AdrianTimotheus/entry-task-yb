package main

import (
	"awesomeProject/db"
	"awesomeProject/redis"
	"awesomeProject/server"
	"fmt"
)

func main() {

	err := db.Connect()
	if err != nil {
		fmt.Println("Error connecting to database")
	}

	err = redis.InitRedis()
	if err != nil {
		fmt.Println("Error connecting to redis")
	}

	server.Init()
}
