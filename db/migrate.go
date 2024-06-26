package db

import (
	"log"

	"github.com/labstack/echo/v4"
)

func Migration(c echo.Context) error {

	userRecord := `CREATE TABLE IF NOT EXISTS UserRecord(
	id serial PRIMARY KEY,
	username VARCHAR(20) NOT NULL,
	userpassword VARCHAR(20) NOT NULL,
	useremail VARCHAR(50) NOT NULL
	);`

	_, err := dbConn.Exec(userRecord)
	if err != nil {
		log.Fatalf("Failed to create UserRecord table: %v", err)
	}

	calculator := `CREATE TABLE IF NOT EXISTS Calculator(
	id serial PRIMARY KEY,
	no1 FLOAT NOT NULL,
	no2 FLOAT NOT NULL CHECK(no2!=0),
	opertion CHAR(1) NOT NULL,
	result FLOAT NOT NULL
	);`

	_, err = dbConn.Exec(calculator)
	if err != nil {
		log.Fatalf("Failed to create Calculator table: %v", err)
	}

	fileStats := `CREATE TABLE IF NOT EXISTS filestats(
	id serial PRIMARY KEY,
	totalline BIGINT,
	totalwords BIGINT,
	totalspaces BIGINT,
	totalvowels BIGINT,
	totalpunctuations BIGINT,
	timestamp timestamp default current_timestamp
	);`

	_, err = dbConn.Exec(fileStats)
	if err != nil {
		log.Fatalf("Failed to create filestats table: %v", err)
	}

	return nil
}
