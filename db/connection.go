package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var Dbcon *sqlx.DB

const (
	DB_Host = "tcp(localhost:3306)"
	DB_User = "root"
	DB_Name = "my_db"
	DB_Pass = "123456"
)

func Connection() {
	var err error

	connect := DB_User + ":" + DB_Pass + "@" + DB_Host + "/" + DB_Name
	Dbcon, err = sqlx.Open("mysql", connect)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	} else {
		fmt.Println("db is connected")
	}
	//defer Dbcon.Close()
	err = Dbcon.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
}
