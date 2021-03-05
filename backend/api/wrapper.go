package api

import (
	"github.com/gorilla/mux"
)

// Wrapper type
type Wrapper struct {
	server *mux.Router
}
