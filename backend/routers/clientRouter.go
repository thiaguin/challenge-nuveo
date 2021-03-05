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

	router.Router.HandleFunc("/client/{id}", controller.GetById).Methods("GET")
	router.Router.HandleFunc("/client", controller.GetAll).Methods("GET")
	router.Router.HandleFunc("/client", controller.Create).Methods("POST")
	router.Router.HandleFunc("/client/{id}", controller.Update).Methods("PUT")
	router.Router.HandleFunc("/client/{id}", controller.Delete).Methods("DELETE")

	return router
}
