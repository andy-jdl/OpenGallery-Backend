package main

import (
	"api/core/config"
	"api/core/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Something wrong with Config Loader")
	}

	service := server.NewServer(cfg)
	log.Fatal(service.Start())

}
