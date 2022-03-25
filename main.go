package main

import (
	"log"

	_config "payment/config"
	_server "payment/server"
)

func main() {
	// Ambil fungsi config Isa
	config := _config.LoadConfig(".")

	// ambil fungsi config Isa
	db := _config.InitDB(config)
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	server := _server.NewServer(&config, db)
	err := server.Run()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
