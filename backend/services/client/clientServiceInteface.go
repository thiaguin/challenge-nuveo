package services

import (
	errors "backend/error"
	"backend/models"
	"io"
)

// ClientServiceInterface interface
type ClientServiceInterface interface {
	GetAll() ([]models.Client, *errors.HTTPError)
	GetById(id string) (*models.Client, *errors.HTTPError)
	Update(id string, body io.Reader) (*models.Client, *errors.HTTPError)
	Create(body io.Reader) (*models.Client, *errors.HTTPError)
	Delete(id string) *errors.HTTPError
}
