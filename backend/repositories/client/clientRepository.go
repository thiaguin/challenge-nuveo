package repositories

import (
	"backend/models"
)

type clientRepository struct{}

// NewClientRepository func
func NewClientRepository() ClientRepositoryInterface {
	return &clientRepository{}
}

func (r clientRepository) GetAll() ([]models.Client, error) {
	galeano := models.Client{Name: "Galeano", Address: "Cotia"}
	nestor := models.Client{Name: "Nestor", Address: "Morumbi"}

	clients := []models.Client{
		galeano,
		nestor,
	}

	return clients, nil
}

func (r clientRepository) GetById(id string) (models.Client, error) {
	luciano := models.Client{Name: "Luciano", Address: "Cotia"}

	return luciano, nil
}

func (r clientRepository) Create(client models.Client) (models.Client, error) {
	return client, nil
}

func (r clientRepository) Update(id string, client models.Client) (models.Client, error) {
	return client, nil
}

func (r clientRepository) Delete(id string) error {
	return nil
}
