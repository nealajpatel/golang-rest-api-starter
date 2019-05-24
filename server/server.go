package server

import "golang-rest-api-starter/config"

func Init() {
	config := config.GetConfig()
	r := SetupRouter()
	r.Run(config.GetString("server.port"))
}
