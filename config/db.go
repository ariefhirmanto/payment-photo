package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(config MainConfig) *sql.DB { //create connection to database
	db, err := sql.Open("mysql", config.Database.DBUser+":"+config.Database.DBPass+
		"@tcp("+config.Database.Host+":"+config.Database.Port+")/"+config.Database.DBName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
