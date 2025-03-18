package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/redscaresu/simpleAPI/models"
)

var notFound = errors.New("no found")

type CityRepository struct {
	cities map[uuid.UUID]models.City
}

func New() *CityRepository {
	return &CityRepository{
		cities: make(map[uuid.UUID]models.City),
	}
}

func (c *CityRepository) AddCity(city *models.City) (*models.City, error) {
	id := uuid.New()
	c.cities[id] = *city
	return city, nil
}

func (c *CityRepository) GetCity(uuid uuid.UUID) (*models.City, error) {
	city, ok := c.cities[uuid]
	if !ok {
		return nil, notFound
	}
	return &city, nil
}
