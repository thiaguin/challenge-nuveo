package tests

import (
	"backend/models"

	"github.com/stretchr/testify/mock"
)

// ClientRepositoryMock type
type ClientRepositoryMock struct {
	mock.Mock
}

func (m *ClientRepositoryMock) GetAll() ([]models.Client, error) {
	args := m.Called()
	clients := []models.Client{}

	return clients, args.Error(1)
}

func (m *ClientRepositoryMock) GetById(id string) (*models.Client, error) {
	args := m.Called()
	client := models.Client{}

	return &client, args.Error(1)
}

func (m *ClientRepositoryMock) Create(client models.Client) (*models.Client, error) {
	args := m.Called()
	return &client, args.Error(1)
}

func (m *ClientRepositoryMock) Update(client *models.Client, value map[string]interface{}) (*models.Client, error) {
	args := m.Called()
	return client, args.Error(1)
}

func (m *ClientRepositoryMock) Delete(client *models.Client) error {
	args := m.Called()
	return args.Error(0)
}
