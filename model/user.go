package model

import (
	"backendmod/config"
	"backendmod/entity"
	"database/sql"
)

type User interface {
	Create(email string, password []byte) (sql.Result, bool)
	GetByEmail(email string) (entity.User, bool)
	SET(email string, dataType string, value string) (sql.Result, bool)
}

type user struct{}

func NewUser() User {
	return &user{}
}

func (*user) Create(email string, password []byte) (sql.Result, bool) {
	db := config.DBConnection()

	full_name := ""
	telephone := ""
	id_google := "0"

	queryString := "INSERT INTO `" + config.DB_NAME + "`.`" + config.DB_USERS_TABLE + "` (`email`, `password`, `full_name`, `telephone`, `id_google`) VALUES (?, ?, ?, ?, ?);"
	stmt, err := db.Prepare(queryString)
	if err != nil {
		return nil, false
	}
	defer db.Close()

	var res sql.Result
	res, err = stmt.Exec(email, password, full_name, telephone, id_google)
	rowsAff, _ := res.RowsAffected()
	if err != nil || rowsAff != 1 {
		return nil, false
	}

	return res, true
}

func (*user) GetByEmail(email string) (entity.User, bool) {
	db := config.DBConnection()
	var user entity.User

	queryString := "SELECT * FROM " + config.DB_NAME + "." + config.DB_USERS_TABLE + " WHERE email = ?;"
	row := db.QueryRow(queryString, email)
	err := row.Scan(
		&user.ID_User,
		&user.Email,
		&user.Password,
		&user.Full_Name,
		&user.Telephone,
		&user.Google,
	)
	if err != nil {
		return user, false
	}

	defer db.Close()
	return user, true
}

func (*user) SET(email string, value string, dataType string) (sql.Result, bool) {
	db := config.DBConnection()

	if dataType != "email" && dataType != "full_name" && dataType != "telephone" {
		return nil, false
	}

	queryString := "UPDATE `" + config.DB_NAME + "`.`" + config.DB_USERS_TABLE + "` SET `" + dataType + "` = ? WHERE (`email` = ?);"

	stmt, err := db.Prepare(queryString)
	if err != nil {
		return nil, false
	}
	defer db.Close()

	var res sql.Result

	res, err = stmt.Exec(value, email)

	rowsAff, err := res.RowsAffected()
	if err != nil || rowsAff != 1 {
		return nil, false
	}

	return res, true
}
