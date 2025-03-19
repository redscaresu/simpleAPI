package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/redscaresu/simpleAPI/models"
	"github.com/redscaresu/simpleAPI/repository"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestGetWeatherCityHandler(t *testing.T) {
	repo := repository.New()
	app := NewApplication(repo)

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

	var info models.Info
	err = json.Unmarshal(body, &info)

	assert.Equal(t, "Dallas", info.Data[0].Name)

	req = httptest.NewRequest(http.MethodGet, "/weather/city?name=Dallas", nil)
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

	assert.Equal(t, city.Name, "Dallas")

}
