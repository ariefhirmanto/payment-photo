package main

import (
	"log"
	_config "payment/config"
	_server "payment/server"
)

func main() {
	// Load config data
	config := _config.LoadConfig(".")

	// Return database connection instance
	db := _config.InitDB(config)
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// insert, err := db.Query("insert into visitor values (4,'goblok','awkoawko@anjay.com',current_timestamp,'FAILED','QRIS');")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()

	server := _server.NewServer(&config, db)
	err := server.Run()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
