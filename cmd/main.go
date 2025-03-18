package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/redscaresu/simpleAPI/handlers"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	handlers.RegisterRoutes(r)
	return http.ListenAndServe(":3000", r)
}
