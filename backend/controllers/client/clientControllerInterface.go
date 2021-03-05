package controllers

import (
	"net/http"
)

// ClientInterface interface
type ClientInterface interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}
