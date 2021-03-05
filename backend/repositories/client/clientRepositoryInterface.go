package repositories

import (
	"backend/models"
)

// ClientRepositoryInterface interface
type ClientRepositoryInterface interface {
	GetAll() ([]models.Client, error)
	GetById(id string) (models.Client, error)
	Create(client models.Client) (models.Client, error)
	Update(id string, client models.Client) (models.Client, error)
	Delete(id string) error
}
