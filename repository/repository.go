package repository

import (
	"errors"

	"github.com/redscaresu/simpleAPI/models"
)

var NotFound = errors.New("not found")

type Repository struct {
	cities map[string]models.City
}

func New() *Repository {
	return &Repository{
		cities: make(map[string]models.City),
	}
}

func (c *Repository) AddCity(city *models.City) *models.City {
	c.cities[city.Name] = *city
	return city
}

func (c *Repository) GetCity(name string) (*models.City, error) {
	if city, ok := c.cities[name]; ok {
		return &city, nil
	}
	return nil, NotFound
}
