package api

import (
	clientController "backend/controllers/client"
	clientRepository "backend/repositories/client"
	routers "backend/routers"
	clientService "backend/services/client"
)

// NewClientServer func
func NewClientServer() *Wrapper {
	clientRepository := clientRepository.NewClientRepository()
	clientService := clientService.NewClientService(clientRepository)
	clientController := clientController.NewClientController(clientService)
	clientRouter := routers.NewClientRouter(clientController)

	return &Wrapper{
		server: clientRouter.Router,
	}
}
