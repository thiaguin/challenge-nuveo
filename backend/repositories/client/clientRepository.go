package repositories

type clientRepository struct{}

// NewClientRepository func
func NewClientRepository() ClientRepositoryInterface {
	return &clientRepository{}
}

func (r clientRepository) GetAll() []string {
	clients := []string{"Galeano", "Bruno Rodrigues", "Pablo", "Luciano"}
	return clients
}
