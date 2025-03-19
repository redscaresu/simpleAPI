package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/redscaresu/simpleAPI/models"
)

type APIClient struct {
	Client http.Client
	URL    string
}

func New(client *http.Client, url string) *APIClient {
	return &APIClient{
		Client: *client,
		URL:    url,
	}
}

func (c *APIClient) Get(url string) (*models.Info, error) {
	resp, err := http.Get(url)
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
