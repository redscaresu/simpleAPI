package main

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

func GetWeather(place string) (*Info, error) {
	path := fmt.Sprintf(url+"/api/weather/search?name=%s", place)
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var info Info
	err = json.Unmarshal(bodyBytes, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func GetWeatherHandler(w http.ResponseWriter, r *http.Request) {

	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "missing param", http.StatusBadRequest)
	}

	info, err := GetWeather(city)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}

	jsonInfo, err := json.Marshal(info)
	if err != nil {
		http.Error(w, "error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInfo)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", GetWeatherHandler)
	http.ListenAndServe(":3333", r)
}
