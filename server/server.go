package server

import (
	"golang-rest-api-starter/config"
	"log"
)

func Init() {
	config := config.GetConfig()
	r := SetupRouter()
	err := r.Run(config.GetString("server.port"))
	if err != nil {
		log.Fatal(err)
	}
}
