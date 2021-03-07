package services

import (
	errors "backend/error"
	"backend/models"
	dto "backend/models/dto"
)

// ClientServiceInterface interface
type ClientServiceInterface interface {
	GetAll() ([]models.Client, *errors.HTTPError)
	GetById(id string) (*models.Client, *errors.HTTPError)
	Create(body dto.CreateClientDTO) (*models.Client, *errors.HTTPError)
	Update(id string, data map[string]interface{}) (*models.Client, *errors.HTTPError)
	Delete(id string) *errors.HTTPError
}
