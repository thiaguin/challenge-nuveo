package api

import (
	"backend/repositories"
	messageService "backend/services/message"
	"net/http"
)

// Server struct
type Server struct{}

// NewServer func
func NewServer() *Server {
	return &Server{}
}

// Run func
func (s Server) Run(port string) {
	messageService := messageService.NewMessageService()
	dbRepository := repositories.NewDB()
	clientServer := NewClientServer(dbRepository, messageService)
	http.Handle("/", clientServer.server)
	http.ListenAndServe(port, nil)
}
