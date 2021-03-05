package services

import (
	clientRepository "backend/repositories/client"
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

func (s clientService) GetAll() []string {
	return s.repository.GetAll()
}
