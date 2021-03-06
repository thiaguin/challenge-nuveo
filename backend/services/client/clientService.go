package services

import (
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

func (s clientService) GetAll() ([]models.Client, error) {
	return s.repository.GetAll()
}

func (s clientService) GetById(id string) (*models.Client, error) {
	return s.repository.GetById(id)
}

func (s clientService) Create(body io.Reader) (*models.Client, error) {
	client := dto.CreateClientDTO{}
	err := json.NewDecoder(body).Decode(&client)

	if err != nil {
		return &models.Client{}, errors.New("Error")
	}

	uuidValue, uuidErr := uuid.NewUUID()

	if uuidErr != nil {
		return nil, uuidErr
	}

	newClient := models.Client{
		ID:      uuidValue,
		Name:    client.Name,
		Address: client.Address,
	}

	return s.repository.Create(newClient)
}

func (s clientService) Update(id string, body io.Reader) (*models.Client, error) {
	client, clientErr := s.repository.GetById(id)

	if clientErr != nil {
		return nil, clientErr
	}

	valuesToUpdate := map[string]interface{}{}

	decodeErr := json.NewDecoder(body).Decode(&valuesToUpdate)

	if decodeErr != nil {
		return nil, decodeErr
	}

	return s.repository.Update(client, valuesToUpdate)
}

func (s clientService) Delete(id string) error {
	client, clientErr := s.repository.GetById(id)

	if clientErr != nil {
		return clientErr
	}

	return s.repository.Delete(client)
}
