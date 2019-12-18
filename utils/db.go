package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

func AddSongToRequestsTable(title string, artist string) (int64, error) {
	log.Print("Request to add to song table received")
	db := establishConnection()
	defer db.Close()
	query := "INSERT INTO songs (title, artist) VALUES ($1, $2)"
	result, err := db.Exec(query, title, artist)
	if err != nil {
		panic(err)
	}
	log.Print("db entry added successfully")
	status, err := result.LastInsertId()
	if err != nil {
		status, err = result.RowsAffected();
	}
	return status, err
}

// shorthand for opening the db connection, remember to close it as this function does not do that.
func establishConnection() *sql.DB {
	host, port, user, password, dbname := getDbVars()
	dbInfoString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbInfoString)
	if err != nil {
		panic(err)
	}
	return db
}

func getDbVars() (interface{}, interface{}, interface{}, interface{}, interface{}) {
	host := os.Getenv("NOCREDIT_DB_HOST")
	port, err := strconv.Atoi(os.Getenv("NOCREDIT_DB_PORT"))
	if err != nil {
		log.Print("Fatal error: " + err.Error())
		panic("Caused by invalid port: " + string(port))
	}
	user := os.Getenv("NOCREDIT_DB_USERNAME")
	password := os.Getenv("NOCREDIT_DB_PW")
	dbname := os.Getenv("NOCREDIT_DB_NAME")
	return host, port, user, password, dbname
}
