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
	if city, ok := c.cities[uuid]; ok {
		return &city, nil
	}
	return nil, notFound
}
