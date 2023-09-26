package main

import (
	"L0/src/db"
)

func main() {
	database := db.Connect()
	err := database.Ping()

	if err != nil {
		panic(err)
	}

	//defer database.Close()
}
