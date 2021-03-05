package api

import (
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
	clientServer := NewClientServer()
	http.Handle("/", clientServer.server)
	http.ListenAndServe(port, nil)
}
