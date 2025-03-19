package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/redscaresu/simpleAPI/client"
	"github.com/redscaresu/simpleAPI/models"
	"github.com/redscaresu/simpleAPI/repository"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestGetWeatherCityHandler(t *testing.T) {

	want := &models.Info{
		Page:       1,
		PerPage:    2,
		Total:      3,
		TotalPages: 4,
		Data: []models.City{
			{
				Name:    "Malvern",
				Status:  []string{"raining"},
				Weather: "12 c",
			},
		},
	}

	bytesInfo, err := json.Marshal(want)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(bytesInfo))
	}))
	defer srv.Close()

	repo := repository.New()
	apiClient := client.New(srv.Client(), srv.URL)

	app := NewApplication(repo, apiClient)

	r := chi.NewRouter()
	app.RegisterRoutes(r)

	req := httptest.NewRequest(http.MethodGet, "/weather/info?city=dallas", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, resp.StatusCode, http.StatusOK)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var got models.Info
	err = json.Unmarshal(body, &got)

	assert.Equal(t, "Malvern", got.Data[0].Name)

	req = httptest.NewRequest(http.MethodGet, "/weather/city?name=Malvern", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp = w.Result()
	defer resp.Body.Close()

	assert.Equal(t, resp.StatusCode, http.StatusOK)

	body, err = io.ReadAll(resp.Body)
	require.NoError(t, err)

	var city models.City
	err = json.Unmarshal(body, &city)
	require.NoError(t, err)

	assert.Equal(t, "Malvern", city.Name)

}
