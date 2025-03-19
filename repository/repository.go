package repository

import (
	"errors"

	"github.com/redscaresu/simpleAPI/models"
)

var notFound = errors.New("no found")

type CityRepository struct {
	cities map[string]models.City
}

func New() *CityRepository {
	return &CityRepository{
		cities: make(map[string]models.City),
	}
}

func (c *CityRepository) AddCity(city *models.City) *models.City {
	c.cities[city.Name] = *city
	return city
}

func (c *CityRepository) GetCity(name string) (*models.City, error) {
	if city, ok := c.cities[name]; ok {
		return &city, nil
	}
	return nil, notFound
}
