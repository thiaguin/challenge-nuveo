package services

import (
	errors "backend/error"
	"backend/models"
	dto "backend/models/dto"
	"io"
)

// ClientServiceInterface interface
type ClientServiceInterface interface {
	GetAll() ([]models.Client, *errors.HTTPError)
	GetById(id string) (*models.Client, *errors.HTTPError)
	Create(body dto.CreateClientDTO) (*models.Client, *errors.HTTPError)
	Update(id string, body io.Reader) (*models.Client, *errors.HTTPError)
	Delete(id string) *errors.HTTPError
}
