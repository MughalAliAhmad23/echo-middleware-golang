package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var dbConn *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "new_password"
	dbname   = "aliahmaddb"
)

func Connect() (*sql.DB, error) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	dbConn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error connecting to database", err)
		return nil, err
	}
	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}
