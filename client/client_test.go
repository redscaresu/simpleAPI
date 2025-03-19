package client_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/redscaresu/simpleAPI/client"
	"github.com/redscaresu/simpleAPI/models"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestGet(t *testing.T) {
	info := &models.Info{
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

	infoBytes, err := json.Marshal(info)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(infoBytes))
	}))
	defer srv.Close()

	apiClient := client.New(srv.Client(), srv.URL)

	got, err := apiClient.Get(srv.URL)
	require.NoError(t, err)

	assert.Equal(t, info.Data[0].Name, got.Data[0].Name)

}
