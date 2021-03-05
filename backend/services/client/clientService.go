package services

import (
	"backend/models"
	dto "backend/models/dto"
	clientRepository "backend/repositories/client"
	"encoding/json"
	"errors"
	"io"
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

func (s clientService) GetById(id string) (models.Client, error) {
	return s.repository.GetById(id)
}

func (s clientService) Create(body io.Reader) (models.Client, error) {
	client := dto.CreateClientDTO{}
	err := json.NewDecoder(body).Decode(&client)

	if err != nil {
		return models.Client{}, errors.New("Error")
	}

	newClient := models.Client{
		Name:    client.Name,
		Address: client.Address,
	}

	return s.repository.Create(newClient)
}

func (s clientService) Update(id string, body io.Reader) error {
	client := models.Client{}
	err := json.NewDecoder(body).Decode(&client)

	if err != nil {
		return errors.New("Error")
	}

	s.repository.Update(id, client)
	return nil
}

func (s clientService) Delete(id string) error {
	return s.repository.Delete(id)
}
