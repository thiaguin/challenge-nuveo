package controllers

import (
	clientService "backend/services/client"
	"encoding/json"
	"net/http"
)

// ClientController type
type clientController struct {
	service clientService.ClientServiceInterface
}

// NewClientController func
func NewClientController(service clientService.ClientServiceInterface) ClientInterface {
	return &clientController{
		service: service,
	}
}

// GetAll func
func (c clientController) GetAll(w http.ResponseWriter, r *http.Request) {
	clients := c.service.GetAll()
	json.NewEncoder(w).Encode(clients)
}
