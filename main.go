package main

import (
	"api/core/config"
	"api/core/repository"
	"api/core/server"
	"log"
)

// TODO (No order)
// Containerize your project (semi done)
// Fix endpoints (done)
// Standardize endpoint data (done)
// Dynamic calls to services (done)
// Understanding your cache and what it can hold.
// (holds locally not like a redis cache that is held across server, done)
// Fix favorites and maybe make faster (this is at the very end tho)

// After top is done
// Fix front end

// After that
// Deploy and test

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Something wrong with Config Loader")
	}
	db, err := repository.NewClient()
	if err != nil {
		panic("Something wrong with DBClient")
	}

	err = db.DBMigrate()
	if err != nil {
		log.Fatal("repository migration failed")
		return
	}

	service := server.NewServer(db, cfg)
	log.Fatal(service.Start())

}
