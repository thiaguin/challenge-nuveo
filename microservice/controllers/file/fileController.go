package controllers

import (
	"fmt"
	file "microservice/services/file"
	message "microservice/services/message"
	"microservice/utils"
	"net/http"
)

type fileController struct {
	messageService message.MessageServiceInterface
	fileService    file.FileServiceInterface
}

// NewFileController func
func NewFileController(messageService message.MessageServiceInterface, fileService file.FileServiceInterface) FileControllerInterface {
	return &fileController{
		messageService: messageService,
		fileService:    fileService,
	}
}

// Write func
func (c fileController) Write(w http.ResponseWriter, r *http.Request) {
	directory := utils.GetEnvVariable("NOVOS_CLIENTES")
	dockerVolume := utils.GetEnvVariable("DOCKER_VOLUME")

	path := fmt.Sprintf("%s%s", dockerVolume, directory)

	_, pathErr := c.fileService.Exist(path)

	if pathErr != nil {
		http.Error(w, "Caminho não encontrado", 400)
		return
	}

	message, messageErr := c.messageService.Dequeue()

	if messageErr != nil {
		http.Error(w, messageErr.Error(), 500)
		return
	}

	if message == nil {
		http.Error(w, "Não tem clientes na fila", 404)
		return
	}

	writeErr := c.fileService.Write(message, path)

	if writeErr != nil {
		http.Error(w, writeErr.Error(), 500)
		return
	}

	w.WriteHeader(200)
}
