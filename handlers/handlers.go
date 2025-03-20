package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/redscaresu/simpleAPI/client"
	"github.com/redscaresu/simpleAPI/repository"
)

type requestWeather struct {
	Name string `json:"name"`
}

func NewApplication(repo *repository.Repository, client *client.APIClient) *application {
	return &application{
		repo:   repo,
		client: client,
	}
}

type application struct {
	repo   *repository.Repository
	client *client.APIClient
}

func (a *application) RegisterRoutes(mux *chi.Mux) {
	mux.Route("/weather", func(r chi.Router) {
		r.Get("/info", a.GetWeatherHandler)
		r.Get("/city", a.GetCityHandler)
		r.Post("/postcity", a.PostCityHandler)
	})
}

func (a *application) GetWeatherHandler(w http.ResponseWriter, r *http.Request) {

	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "missing param", http.StatusBadRequest)
		return
	}

	info, err := a.client.Get(fmt.Sprintf(a.client.URL+"/api/weather/search?name=%s", city))
	if err != nil {
		http.Error(w, fmt.Sprintf("internal error: %d", err), http.StatusInternalServerError)
		return
	}

	for _, v := range info.Data {
		a.repo.AddCity(&v)
	}

	jsonInfo, err := json.Marshal(info)
	if err != nil {
		http.Error(w, "error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInfo)
}

func (a *application) GetCityHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	city, err := a.repo.GetCity(name)
	if err != nil {
		http.Error(w, "city not found", http.StatusNotFound)
		return
	}

	jsonCity, err := json.Marshal(city)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonCity)
}

func (a *application) PostCityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "no body found", http.StatusBadRequest)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read body", http.StatusBadRequest)
		return
	}

	var reqWeather requestWeather
	err = json.Unmarshal(bytes, &reqWeather)

	c, err := a.repo.GetCity(reqWeather.Name)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	cBytes, err := json.Marshal(c)

	w.Header().Set("Content-Type", "application/json")
	w.Write(cBytes)
}
