package routers

import (
	clientController "backend/controllers/client"

	"github.com/gorilla/mux"
)

// NewClientRouter func
func NewClientRouter(controller clientController.ClientInterface) *Router {
	router := &Router{
		Router: mux.NewRouter(),
	}

	router.Router.HandleFunc("/client", controller.GetAll).Methods("GET")

	return router
}
