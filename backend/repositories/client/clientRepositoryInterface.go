package repositories

import (
	"backend/models"
)

// ClientRepositoryInterface interface
type ClientRepositoryInterface interface {
	GetAll() ([]models.Client, error)
	GetById(id string) (*models.Client, error)
	Create(client models.Client) (*models.Client, error)
	Update(client *models.Client, value map[string]interface{}) (*models.Client, error)
	Delete(client *models.Client) error
}
