package main

import (
	"log"
	_config "payment/config"
	_server "payment/server"
)

func main() {
	// Load config data
	config := _config.LoadConfig()

	// Return database connection instance
	db := _config.InitDB(config)

	server := _server.NewServer(&config, db)
	err := server.Run()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
