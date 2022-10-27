package config

import (
	"database/sql"
	"log"
)

func DBConnection() (db *sql.DB) {
	db, err := sql.Open(DB_DRIVER, DB_USER+":"+DB_PASS+"@"+DB_PROTOCOL+"("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)

	if err != nil {
		log.Println("Error validating sql.Open arguments.")
		log.Println(err.Error())
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Println("Error verifying connection with db.Ping.")
		log.Println(err.Error())
		return nil
	}

	return db
}
