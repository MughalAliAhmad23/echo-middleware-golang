package db

import (
	"calculator/models"
	"fmt"
)

func Insert(data models.CalculatorDb) (*models.CalculatorDb, error) {
	query := "INSERT INTO Calculator (no1, no2, opertion, result) VALUES ($1,$2,$3,$4)RETURNING id"

	err := dbConn.QueryRow(query, data.No1, data.No2, data.Operation, data.Result).Scan(&data.Id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &data, nil
}

func UserInsert(user models.UserDb) (*models.UserDb, error) {
	query := "INSERT INTO UserRecord (username,userpassword,useremail) VALUES ($1,$2,$3) RETURNING id"

	err := dbConn.QueryRow(query, user.Username, user.Userpassword, user.Useremail).Scan(&user.Id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func Fileinsert(stats models.FilestatsDB) (*models.FilestatsDB, error) {
	query := "INSERT INTO filestats (totalline,totalwords,totalspaces,totalvowels,totalpunctuations) VALUES ($1,$2,$3,$4,$5) RETURNING id, timestamp"

	err := dbConn.QueryRow(query, stats.Totallines, stats.Totalwords, stats.Totalspaces, stats.Totalvowels, stats.Totalpunctuation).Scan(&stats.Id, &stats.Timestamp)

	if err != nil {
		return nil, err
	}

	return &stats, nil
}

func Userlogin(user models.Userlogin) (bool, error) {
	query := "SELECT 'exists' AS result FROM UserRecord WHERE username = $1 AND userpassword = $2 UNION SELECT 'not exists' AS result LIMIT 1"

	rows, err := dbConn.Query(query, user.Username, user.Userpassword)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	found := false
	for rows.Next() {
		var result string
		err := rows.Scan(&result)
		if err != nil {
			return false, err
		}
		if result == "exists" {
			found = true
		}
	}
	err = rows.Err()
	if err != nil {
		return false, err
	}
	return found, nil
}

func Isuserexists(user models.User) (bool, error) {
	isuser := false
	query := "SELECT CASE  WHEN EXISTS (SELECT 1 FROM UserRecord WHERE username = $1) THEN 'exists' ELSE 'notexists' END AS result"

	rows, err := dbConn.Query(query, user.Username)
	if err != nil {
		return false, nil
	}
	defer rows.Close()
	for rows.Next() {
		var result string
		err := rows.Scan(&result)
		if err != nil {
			return false, nil
		}
		if result == "exists" {
			isuser = true
		}
	}
	err = rows.Err()
	if err != nil {
		return false, nil
	}
	return isuser, nil
}

func Readall() ([]models.CalculatorDb, error) {

	rows, err := dbConn.Query("SELECT * FROM Calculator")
	if err != nil {
		return nil, err
	}
	var calculations []models.CalculatorDb
	for rows.Next() {
		var cal models.CalculatorDb
		err := rows.Scan(&cal.Id, &cal.No1, &cal.No2, &cal.Operation, &cal.Result)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		calculations = append(calculations, cal)
	}
	return calculations, nil
}
func Readallfilestats() ([]models.FilestatsDB, error) {
	rows, err := dbConn.Query("SELECT * FROM filestats")
	if err != nil {
		return nil, err
	}
	var filestatsdb []models.FilestatsDB
	for rows.Next() {
		var filestats models.FilestatsDB
		err := rows.Scan(&filestats.Id, &filestats.Totallines, &filestats.Totalwords, &filestats.Totalspaces, &filestats.Totalvowels, &filestats.Totalpunctuation, &filestats.Timestamp)
		if err != nil {
			return nil, err
		}
		filestatsdb = append(filestatsdb, filestats)
	}
	return filestatsdb, nil
}
func Readbyid(id int) (*models.CalculatorDb, error) {
	var cal models.CalculatorDb

	err := dbConn.QueryRow("SELECT * FROM Calculator WHERE id =$1", id).Scan(&cal.Id, &cal.No1, &cal.No2, &cal.Operation, &cal.Result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &cal, nil
}

func Readbysymbol(sym string) ([]models.CalculatorDb, error) {
	fmt.Println(sym)

	rows, err := dbConn.Query("SELECT * FROM Calculator WHERE opertion=$1", sym)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var calculation []models.CalculatorDb
	for rows.Next() {
		var cal models.CalculatorDb
		err := rows.Scan(&cal.Id, &cal.No1, &cal.No2, &cal.Operation, &cal.Result)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		calculation = append(calculation, cal)
	}
	return calculation, nil
}
func Removebyid(id int) error {

	_, err := dbConn.Exec("DELETE FROM Calculator Where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func Reupdate(id int, no1, no2, res float64) error {

	_, err := dbConn.Exec("UPDATE Calculator SET no1=$1, no2=$2, result=$3 WHERE id=$4", no1, no2, res, id)
	if err != nil {
		return err
	}

	return err
}
