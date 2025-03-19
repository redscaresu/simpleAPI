package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/redscaresu/simpleAPI/client"
	"github.com/redscaresu/simpleAPI/handlers"
	"github.com/redscaresu/simpleAPI/repository"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	repo := repository.New()
	client := client.New(httpClient, "https://jsonmock.hackerrank.com")

	app := handlers.NewApplication(repo, client)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	app.RegisterRoutes(r)
	return http.ListenAndServe(":3000", r)
}
