package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type clientRepository struct {
	db *gorm.DB
}

// NewClientRepository func
func NewClientRepository(db *gorm.DB) ClientRepositoryInterface {
	return &clientRepository{
		db: db,
	}
}

func (r clientRepository) GetAll() ([]models.Client, error) {
	clients := []models.Client{}
	result := r.db.Find(&clients)

	if result.Error != nil {
		return nil, result.Error
	}

	return clients, nil
}

func (r clientRepository) GetById(id string) (*models.Client, error) {
	client := models.Client{}
	result := r.db.Where("id = ?", id).First(&client)

	if result.Error != nil {
		return nil, result.Error
	}

	return &client, nil
}

func (r clientRepository) Create(client models.Client) (*models.Client, error) {
	result := r.db.Create(&client)

	if result.Error != nil {
		return nil, result.Error
	}

	return &client, nil
}

func (r clientRepository) Update(client *models.Client, values map[string]interface{}) (*models.Client, error) {
	r.db.Model(client).Omit("id").Updates(values)
	return client, nil
}

func (r clientRepository) Delete(client *models.Client) error {
	result := r.db.Unscoped().Delete(client)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
