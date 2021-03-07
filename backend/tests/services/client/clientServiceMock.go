package tests

import (
	customError "backend/error"
	"backend/models"
	dto "backend/models/dto"

	"github.com/stretchr/testify/mock"
)

// ClientServiceMock type
type ClientServiceMock struct {
	mock.Mock
}

func (m *ClientServiceMock) GetAll() ([]models.Client, *customError.HTTPError) {
	args := m.Called()
	clients := []models.Client{}

	if args.Get(1) == nil {
		return clients, nil
	}

	return clients, &customError.HTTPError{Status: args.Int(1)}
}

func (m *ClientServiceMock) GetById(id string) (*models.Client, *customError.HTTPError) {
	args := m.Called()
	client := models.Client{}

	if args.Get(1) == nil {
		return &client, nil
	}

	return &client, &customError.HTTPError{Status: args.Int(1)}
}

func (m *ClientServiceMock) Create(body dto.CreateClientDTO) (*models.Client, *customError.HTTPError) {
	args := m.Called()

	if args.Get(1) == nil {
		return &models.Client{}, nil
	}

	return &models.Client{}, &customError.HTTPError{Status: args.Int(1)}
}

func (m *ClientServiceMock) Update(id string, data map[string]interface{}) (*models.Client, *customError.HTTPError) {
	args := m.Called()

	if args.Get(1) == nil {
		return &models.Client{}, nil
	}

	return &models.Client{}, &customError.HTTPError{Status: args.Int(1)}
}

func (m *ClientServiceMock) Delete(id string) *customError.HTTPError {
	args := m.Called()

	if args.Get(0) == nil {
		return nil
	}

	return &customError.HTTPError{Status: args.Int(0)}
}
