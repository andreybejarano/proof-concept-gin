package main

import (
	"fmt"
	"proof-concept-gin/config"
	"proof-concept-gin/route"
)

func main() {
	config.Init()
	configs := config.GetConfig()
	router := route.SetupRouter()
	port := fmt.Sprintf(":%s", configs["server"]["port"])
	router.Run(port)
}
