package api

import (
	clientController "backend/controllers/client"
	clientRepository "backend/repositories/client"
	routers "backend/routers"
	clientService "backend/services/client"

	"gorm.io/gorm"
)

// NewClientServer func
func NewClientServer(db *gorm.DB) *Wrapper {
	clientRepository := clientRepository.NewClientRepository(db)
	clientService := clientService.NewClientService(clientRepository)
	clientController := clientController.NewClientController(clientService)
	clientRouter := routers.NewClientRouter(clientController)

	return &Wrapper{
		server: clientRouter.Router,
	}
}
