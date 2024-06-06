package db

import (
	"calculator/models"
	"fmt"
)

func Insert(data models.CalculatorDb) (*models.CalculatorDb, error) {
	query := "INSERT INTO Calculator (no1, no2, operation, result) VALUES (?,?,?,?)"
	stmt, err := Dbcon.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer stmt.Close()
	result, err := Dbcon.Exec(query, data.No1, data.No2, data.Operation, data.Result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	data.Id = int(lastInsertedId)
	return &data, nil
}

func UserInsert(user models.UserDb) (*models.UserDb, error) {
	query := "INSERT INTO UserRecord (username,userpassword,useremail) VALUES (?,?,?)"
	stmt, err := Dbcon.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := Dbcon.Exec(query, user.Username, user.Userpassword, user.Useremail)
	if err != nil {
		return nil, err
	}
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.Id = int(lastInsertedId)
	return &user, nil
}

func Userlogin(user models.Userlogin) (bool, error) {
	query := "SELECT 'exists' AS result FROM UserRecord WHERE username = ? AND userpassword = ? UNION SELECT 'not exists' AS result LIMIT 1"
	rows, err := Dbcon.Query(query, user.Username, user.Userpassword)
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
	query := "SELECT 'exists' AS result FROM UserRecord WHERE username = ? UNION SELECT 'notexists' AS resut LIMIT 1"
	rows, err := Dbcon.Query(query, user.Username)
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
	rows, err := Dbcon.Queryx("SELECT * FROM Calculator")
	if err != nil {
		return nil, err
	}
	var calculations []models.CalculatorDb
	for rows.Next() {
		var cal models.CalculatorDb
		err := rows.StructScan(&cal)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		calculations = append(calculations, cal)
	}
	return calculations, nil
}

func Readbyid(id int) (*models.CalculatorDb, error) {
	var cal models.CalculatorDb
	err := Dbcon.QueryRowx("SELECT * FROM Calculator WHERE id =?", id).StructScan(&cal)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &cal, nil
}

func Readbysymbol(sym string) ([]models.CalculatorDb, error) {
	fmt.Println(sym)
	rows, err := Dbcon.Queryx("SELECT * FROM Calculator WHERE operation=?", sym)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var calculation []models.CalculatorDb
	for rows.Next() {
		var cal models.CalculatorDb
		err := rows.StructScan(&cal)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		calculation = append(calculation, cal)
	}
	return calculation, nil
}
func Removebyid(id int) error {
	_, err := Dbcon.Exec("DELETE FROM Calculator Where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func Reupdate(id int, no1, no2, res float64) error {
	_, err := Dbcon.Exec("UPDATE Calculator SET no1=?, no2=?, result=? WHERE id=?", no1, no2, res, id)
	if err != nil {
		return err
	}

	return err
}
