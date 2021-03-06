package services

import (
	customError "backend/error"
	"backend/models"
	dto "backend/models/dto"
	clientRepository "backend/repositories/client"
	messageService "backend/services/message"
	"encoding/json"
	"errors"
	"io"

	"github.com/google/uuid"
)

type clientService struct {
	repository     clientRepository.ClientRepositoryInterface
	messageService messageService.MessageServiceInterface
}

// NewClientService func
func NewClientService(repository clientRepository.ClientRepositoryInterface, messageService messageService.MessageServiceInterface) ClientServiceInterface {
	return &clientService{
		repository:     repository,
		messageService: messageService,
	}
}

// GetAll func
func (s clientService) GetAll() ([]models.Client, *customError.HTTPError) {
	clients, err := s.repository.GetAll()

	if err != nil {
		return nil, customError.NewHTTPError(err, 404, "Client")
	}

	return clients, nil
}

// GetById func
func (s clientService) GetById(id string) (*models.Client, *customError.HTTPError) {
	client, err := s.repository.GetById(id)

	if err != nil {
		return nil, customError.NewHTTPError(err, 404, "Client")
	}

	return client, nil
}

// Create func
func (s clientService) Create(body io.Reader) (*models.Client, *customError.HTTPError) {
	client, clientErr := getValidClient(body)

	if clientErr != nil {
		return nil, clientErr
	}

	uuidValue, uuidErr := uuid.NewUUID()

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

	messageErr := enqueueClient(clientCreated, s.messageService)

	if messageErr != nil {
		return nil, messageErr
	}

	return clientCreated, nil
}

// Update func
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

// Delete func
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

func enqueueClient(client *models.Client, messageService messageService.MessageServiceInterface) *customError.HTTPError {
	clientToEnqueue, clientToEnqueueErr := json.Marshal(client)

	if clientToEnqueueErr != nil {
		return customError.NewHTTPError(clientToEnqueueErr, 500, "Encode Client to message")
	}

	messageErr := messageService.Enqueue(clientToEnqueue)

	if messageErr != nil {
		return customError.NewHTTPError(clientToEnqueueErr, 500, "Message Error")
	}

	return nil
}

func getValidClient(body io.Reader) (dto.CreateClientDTO, *customError.HTTPError) {
	client := dto.CreateClientDTO{}
	decodeErr := json.NewDecoder(body).Decode(&client)

	if decodeErr != nil {
		return dto.CreateClientDTO{}, customError.NewHTTPError(decodeErr, 400, "BadRequest")
	}

	if client.Name == "" || client.Address == "" {
		return dto.CreateClientDTO{}, customError.NewHTTPError(errors.New("Validation error"), 400, "BadRequest")
	}

	return client, nil
}
