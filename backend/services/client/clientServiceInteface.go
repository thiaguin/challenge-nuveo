package services

import (
	"backend/models"
	"io"
)

// ClientServiceInterface interface
type ClientServiceInterface interface {
	GetAll() ([]models.Client, error)
	GetById(id string) (models.Client, error)
	Update(id string, body io.Reader) error
	Create(body io.Reader) (models.Client, error)
	Delete(id string) error
}
