package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	url = "https://jsonmock.hackerrank.com"
)

type Info struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []City `json:"data"`
}

type City struct {
	Name    string   `json:"name"`
	Status  []string `json:"status"`
	Weather string   `json:"weather"`
}

func RegisterRoutes(mux *chi.Mux) {
	mux.Route("/weather", func(r chi.Router) {
		r.Get("/", GetWeatherHandler)
	})
}

func GetWeather(place string) (*Info, error) {
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

	var info Info
	err = json.Unmarshal(bodyBytes, &info)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &info, nil
}

func GetWeatherHandler(w http.ResponseWriter, r *http.Request) {

	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "missing param", http.StatusBadRequest)
		return
	}

	info, err := GetWeather(city)
	if err != nil {
		http.Error(w, fmt.Sprintf("internal error: %d", err), http.StatusInternalServerError)
		return
	}

	jsonInfo, err := json.Marshal(info)
	if err != nil {
		http.Error(w, "error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInfo)
}
