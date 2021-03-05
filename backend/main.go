package main

import (
	"backend/api"
)

func main() {
	server := api.NewServer()
	server.Run(":3001")
}
