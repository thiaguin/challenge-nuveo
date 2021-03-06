package main

import (
	"backend/api"
	"backend/utils"
	"fmt"
)

func main() {
	server := api.NewServer()
	server.Run(fmt.Sprintf(":%s", utils.GetEnvVariable("PORT")))
}
