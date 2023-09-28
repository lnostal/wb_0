package api

import (
	v1 "L0/src/api/v1"
	"github.com/go-chi/chi/v5"
)

const prefix = "/api/v1"

func API(r *chi.Mux) {

	r.Group(func(r chi.Router) {

		r.Get(prefix+"/order/{id:[0-9]+}", v1.GetOrderById)

	})
}
