package model

import (
	"backendmod/config"
	"database/sql"
	"log"
)

type GoogleUser interface {
	SET(email string, id_google string) (sql.Result, bool)
	CREATE(email string, name string) (sql.Result, bool)
}

type googleUser struct{}

func NewGoogleUser() GoogleUser {
	return &googleUser{}
}

func (*googleUser) SET(email string, id_google string) (sql.Result, bool) {
	db := config.DBConnection()

	queryString := "UPDATE `" + config.DB_NAME + "`.`" + config.DB_USERS_TABLE + "` SET `id_google` = ? WHERE (`email` = ?);"
	stmt, err := db.Prepare(queryString)
	if err != nil {
		log.Println("Error preparing queryString")
		return nil, false
	}
	defer db.Close()

	var res sql.Result
	res, err = stmt.Exec(id_google, email)
	rowsAff, _ := res.RowsAffected()
	if err != nil || rowsAff != 1 {
		log.Println(err)
		log.Println("There was a problem updating the google_id.")
		return nil, false
	}

	return res, true
}

func (*googleUser) CREATE(email string, name string) (sql.Result, bool) {
	db := config.DBConnection()

	full_name := name
	telephone := ""
	id_google := "1"

	queryString := "INSERT INTO `" + config.DB_NAME + "`.`" + config.DB_USERS_TABLE + "` (`email`, `full_name`, `telephone`, `id_google`) VALUES (?, ?, ?, ?);"
	stmt, err := db.Prepare(queryString)
	if err != nil {
		log.Println("CreateGoogleUser- Error preparing queryString")
		return nil, false
	}
	defer db.Close()

	var res sql.Result
	res, err = stmt.Exec(email, full_name, telephone, id_google)
	rowsAff, _ := res.RowsAffected()
	if err != nil || rowsAff != 1 {
		log.Println(err)
		log.Println("CreateGoogleUser - There was a problem creating google user.")
		return nil, false
	}

	return res, true
}

func SetAsGoogleUser(email string, id_google string) (sql.Result, bool) {
	db := config.DBConnection()

	queryString := "UPDATE `" + config.DB_NAME + "`.`" + config.DB_USERS_TABLE + "` SET `id_google` = ? WHERE (`email` = ?);"
	stmt, err := db.Prepare(queryString)
	if err != nil {
		log.Println("Error preparing queryString")
		return nil, false
	}
	defer db.Close()

	var res sql.Result
	res, err = stmt.Exec(id_google, email)
	rowsAff, _ := res.RowsAffected()
	if err != nil || rowsAff != 1 {
		log.Println(err)
		log.Println("There was a problem updating the google_id.")
		return nil, false
	}

	return res, true
}

func CreateGoogleUser(email string, name string) (sql.Result, bool) {
	db := config.DBConnection()

	full_name := name
	telephone := "Telephone"
	id_google := "1"

	queryString := "INSERT INTO `" + config.DB_NAME + "`.`" + config.DB_USERS_TABLE + "` (`email`, `full_name`, `telephone`, `id_google`) VALUES (?, ?, ?, ?);"
	stmt, err := db.Prepare(queryString)
	if err != nil {
		log.Println("CreateGoogleUser- Error preparing queryString")
		return nil, false
	}
	defer db.Close()

	var res sql.Result
	res, err = stmt.Exec(email, full_name, telephone, id_google)
	rowsAff, _ := res.RowsAffected()
	if err != nil || rowsAff != 1 {
		log.Println(err)
		log.Println("CreateGoogleUser - There was a problem creating google user.")
		return nil, false
	}

	return res, true
}
