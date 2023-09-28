package main

import (
	"L0/src/api"
	"L0/src/cache"
	"L0/src/nats"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	cache.New()

	go nats.Connect()

	router := chi.NewRouter()
	go api.API(router)

	http.ListenAndServe(":3333", router)

}
