package controllers

import "net/http"

// FileControllerInterface interface
type FileControllerInterface interface {
	Write(w http.ResponseWriter, r *http.Request)
}
