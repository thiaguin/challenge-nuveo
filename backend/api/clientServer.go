package api

import (
	clientController "backend/controllers/client"
	clientRepository "backend/repositories/client"
	routers "backend/routers"
	clientService "backend/services/client"
	messageService "backend/services/message"

	"gorm.io/gorm"
)

// NewClientServer func
func NewClientServer(db *gorm.DB, messageService messageService.MessageServiceInterface) *Wrapper {
	clientRepository := clientRepository.NewClientRepository(db)
	clientService := clientService.NewClientService(clientRepository, messageService)
	clientController := clientController.NewClientController(clientService)
	clientRouter := routers.NewClientRouter(clientController)

	return &Wrapper{
		server: clientRouter.Router,
	}
}
