package main

import (
	"L0/src/api"
	"L0/src/db"
	"L0/src/nats"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	database := db.Connect()
	err := database.Ping()

	if err != nil {
		panic(err)
	}

	go nats.Connect()

	addr := ":3333"
	fmt.Printf("Starting server on %v\n", addr)
	router := chi.NewRouter()
	api.API(router)
	http.ListenAndServe(addr, router)

}
