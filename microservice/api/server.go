package api

import (
	fileController "microservice/controllers/file"
	fileService "microservice/services/file"
	messageService "microservice/services/message"
	"net/http"
)

// Server type
type Server struct{}

func NewServer() *Server {
	return &Server{}
}

// Run func
func (s Server) Run(port string) {
	fileService := fileService.NewFileService()
	messageService := messageService.NewMessageService()
	fileController := fileController.NewFileController(messageService, fileService)
	http.HandleFunc("/", fileController.Write)
	http.ListenAndServe(port, nil)
}

func show(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}
