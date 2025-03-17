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

type Handler func(w http.ResponseWriter, r *http.Request) error

func GetWeatherHandler(w http.ResponseWriter, r *http.Request) error {

	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "missing param", http.StatusBadRequest)
	}

	info, err := GetWeather(city)
	if err != nil {
		w.Write([]byte("error"))
	}

	jsonInfo, err := json.Marshal(info)
	if err != nil {
		return nil
	}

	w.Write(jsonInfo)
	return nil
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func main() {
	r := chi.NewRouter()
	r.Method("GET", "/", Handler(GetWeatherHandler))
	http.ListenAndServe(":3333", r)

	GetWeather("Dallas")
}
