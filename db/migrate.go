package db

import (
	"log"

	"github.com/labstack/echo"
)

func Migration(c echo.Context) error {

	defer dbConn.Close()

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

	return nil
}
