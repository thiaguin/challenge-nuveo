package api

import (
	"backend/repositories"
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
	dbRepository := repositories.NewDB()
	clientServer := NewClientServer(dbRepository)
	http.Handle("/", clientServer.server)
	http.ListenAndServe(port, nil)
}
