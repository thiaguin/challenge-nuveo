package controllers

import (
	errors "backend/error"
	clientService "backend/services/client"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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
	clients, err := c.service.GetAll()

	if err != nil {
		errors.SendError(w, "Olha o erro", 500)
		return
	}

	json.NewEncoder(w).Encode(clients)
}

// GetAll func
func (c clientController) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientId := params["id"]
	client, err := c.service.GetById(clientId)

	if err != nil {
		errors.SendError(w, "Olha o erro", 500)
		return
	}

	json.NewEncoder(w).Encode(client)
}

// GetAll func
func (c clientController) Create(w http.ResponseWriter, r *http.Request) {
	client, err := c.service.Create(r.Body)

	if err != nil {
		errors.SendError(w, err, 500)
		return
	}

	json.NewEncoder(w).Encode(client)
}

// GetAll func
func (c clientController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientId := params["id"]
	err := c.service.Update(clientId, r.Body)

	if err != nil {
		errors.SendError(w, "Olha o erro", 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAll func
func (c clientController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientId := params["id"]
	err := c.service.Delete(clientId)

	if err != nil {
		errors.SendError(w, "Olha o erro", 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}