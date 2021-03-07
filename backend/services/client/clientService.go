package services

import (
	customError "backend/error"
	"backend/models"
	dto "backend/models/dto"
	clientRepository "backend/repositories/client"
	messageService "backend/services/message"
	"encoding/json"
	"errors"
	"strings"

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
		return nil, customError.NewHTTPError(err, 500, "InternalError")
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
func (s clientService) Create(client dto.CreateClientDTO) (*models.Client, *customError.HTTPError) {
	if strings.Trim(client.Name, " ") == "" || strings.Trim(client.Address, " ") == "" {
		return nil, customError.NewHTTPError(errors.New("Validation error"), 400, "BadRequest")
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
		return nil, customError.NewHTTPError(clientCreatedErr, 400, "Database")
	}

	messageErr := enqueueClient(clientCreated, s.messageService)

	if messageErr != nil {
		return nil, messageErr
	}

	return clientCreated, nil
}

// Update func
func (s clientService) Update(id string, data map[string]interface{}) (*models.Client, *customError.HTTPError) {
	client, clientErr := s.repository.GetById(id)

	if clientErr != nil {
		return nil, customError.NewHTTPError(clientErr, 404, "Client")
	}

	updatedClient, updatedClientErr := s.repository.Update(client, data)

	if updatedClientErr != nil {
		return nil, customError.NewHTTPError(updatedClientErr, 500, "Database")
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
		return customError.NewHTTPError(deleteErr, 500, "Database")
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
		return customError.NewHTTPError(clientToEnqueueErr, 500, "Message Service")
	}

	return nil
}
