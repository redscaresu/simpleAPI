package repository_test

import (
	"testing"

	"github.com/redscaresu/simpleAPI/models"
	"github.com/redscaresu/simpleAPI/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddCityGetCity(t *testing.T) {

	expectedCity := &models.City{
		Name:    "Dallas",
		Status:  []string{"foo"},
		Weather: "boo",
	}

	repo := repository.New()

	repo.AddCity(expectedCity)
	city, err := repo.GetCity("Dallas")
	require.NoError(t, err)

	require.Equal(t, city, expectedCity)
}

func TestCityNotFound(t *testing.T) {

	repo := repository.New()
	city, err := repo.GetCity("Dallas")
	require.Error(t, err)

	assert.Nil(t, city)
	assert.EqualError(t, repository.NotFound, err.Error())
}
