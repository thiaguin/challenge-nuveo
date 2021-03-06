package services

import (
	customError "backend/error"
	"backend/models"
	dto "backend/models/dto"
	clientRepository "backend/repositories/client"
	"encoding/json"
	"errors"
	"io"

	"github.com/google/uuid"
)

type clientService struct {
	repository clientRepository.ClientRepositoryInterface
}

// NewClientService func
func NewClientService(repository clientRepository.ClientRepositoryInterface) ClientServiceInterface {
	return &clientService{
		repository: repository,
	}
}

func (s clientService) GetAll() ([]models.Client, *customError.HTTPError) {
	clients, err := s.repository.GetAll()

	if err != nil {
		return nil, customError.NewHTTPError(err, 404, "Client")
	}

	return clients, nil
}

func (s clientService) GetById(id string) (*models.Client, *customError.HTTPError) {
	client, err := s.repository.GetById(id)

	if err != nil {
		return nil, customError.NewHTTPError(err, 404, "Client")
	}

	return client, nil
}

func (s clientService) Create(body io.Reader) (*models.Client, *customError.HTTPError) {
	client := dto.CreateClientDTO{}
	uuidValue, uuidErr := uuid.NewUUID()
	decodeErr := json.NewDecoder(body).Decode(&client)

	if decodeErr != nil {
		return &models.Client{}, customError.NewHTTPError(decodeErr, 400, "BadRequest")
	}

	if client.Name == "" || client.Address == "" {
		return &models.Client{}, customError.NewHTTPError(errors.New("Validation error"), 400, "BadRequest")
	}

	if uuidErr != nil {
		return nil, customError.NewHTTPError(uuidErr, 500, "UUID")
	}

	newClient := models.Client{
		ID:      uuidValue,
		Name:    client.Name,
		Address: client.Address,
	}

	clientCreated, clientCreatedErr := s.repository.Create(newClient)

	if clientCreatedErr != nil {
		return nil, customError.NewHTTPError(uuidErr, 400, "UUID")
	}

	return clientCreated, nil
}

func (s clientService) Update(id string, body io.Reader) (*models.Client, *customError.HTTPError) {
	client, clientErr := s.repository.GetById(id)

	if clientErr != nil {
		return nil, customError.NewHTTPError(clientErr, 404, "Client")
	}

	valuesToUpdate := map[string]interface{}{}
	decodeErr := json.NewDecoder(body).Decode(&valuesToUpdate)

	if decodeErr != nil {
		return nil, customError.NewHTTPError(clientErr, 400, "BadRequest")
	}

	updatedClient, updatedClientErr := s.repository.Update(client, valuesToUpdate)

	if updatedClient != nil {
		return nil, customError.NewHTTPError(updatedClientErr, 400, "BadRequest")
	}

	return updatedClient, nil
}

func (s clientService) Delete(id string) *customError.HTTPError {
	client, clientErr := s.repository.GetById(id)

	if clientErr != nil {
		return customError.NewHTTPError(clientErr, 404, "Client")
	}

	deleteErr := s.repository.Delete(client)

	if deleteErr != nil {
		return customError.NewHTTPError(deleteErr, 400, "Delete")
	}

	return nil
}
