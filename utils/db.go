package utils

import (
	"database/sql"
	"fmt"
)

const (
	host     = "HOST"
	port     = "PORT"
	user     = "USER"
	password = "PASSWORD"
	dbname   = "DBNAME"
)
func AddSongToRequestsTable(title string, artist string) (int64, error) {
	db := establishConnection()
	defer db.Close()
	query := "INSERT INTO songs (title, artist) VALUES ($1, $2)"
	result, err := db.Exec(query, title, artist)
	if err != nil {
		panic(err)
	}
	return result.LastInsertId()
}

func establishConnection() *sql.DB {
	dbInfoString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbInfoString)
	if err != nil {
		panic(err)
	}
	return db
}
