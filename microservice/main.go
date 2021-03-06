package main

import (
	"fmt"
	"microservice/api"
	"microservice/utils"
)

func main() {
	server := api.NewServer()
	server.Run(fmt.Sprintf(":%s", utils.GetEnvVariable("PORT")))
}
