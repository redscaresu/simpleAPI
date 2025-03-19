package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/redscaresu/simpleAPI/models"
	"github.com/redscaresu/simpleAPI/repository"
)

const (
	url = "https://jsonmock.hackerrank.com"
)

func NewApplication(repo *repository.Repository) *application {
	return &application{
		repo: repo,
	}
}

type application struct {
	repo *repository.Repository
}

func (a *application) RegisterRoutes(mux *chi.Mux) {
	mux.Route("/weather", func(r chi.Router) {
		r.Get("/info", a.GetWeatherHandler)
		r.Get("/city", a.GetCityHandler)
	})
}

func (a *application) GetWeatherClient(place string) (*models.Info, error) {
	path := fmt.Sprintf(url+"/api/weather/search?name=%s", place)
	resp, err := http.Get(path)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var info models.Info
	err = json.Unmarshal(bodyBytes, &info)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &info, nil
}

func (a *application) GetWeatherHandler(w http.ResponseWriter, r *http.Request) {

	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "missing param", http.StatusBadRequest)
		return
	}

	info, err := a.GetWeatherClient(city)
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
