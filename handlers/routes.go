package handlers

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(mux *chi.Mux) {
	mux.Route("/weather", func(r chi.Router) {
		r.Get("/", GetWeatherHandler)
	})
}
