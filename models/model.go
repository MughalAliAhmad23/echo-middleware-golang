package models

import "time"

type Resp struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

type CalculatorReq struct {
	No1 float64 `json:"no1"`
	No2 float64 `json:"no2"`
}

type CalculatorResp struct {
	Id        int     `json:"id"`
	No1       float64 `json:"no1"`
	No2       float64 `json:"no2"`
	Operation string  `json:"operation"`
	Result    float64 `json:"result"`
}
type CalculatorDb struct {
	Id        int     `db:"id" json:"id"`
	No1       float64 `db:"no1" json:"no1"`
	No2       float64 `db:"no2" json:"no2"`
	Operation string  `db:"operation" json:"operation"`
	Result    float64 `db:"result" json:"result"`
}

type User struct {
	Username     string `json:"Username"`
	Userpassword string `json:"Userpassword"`
	Useremail    string `json:"Useremail"`
}
type UserDb struct {
	Id           int    `db:"id"`
	Username     string `db:"username"`
	Userpassword string `db:"userpassword"`
	Useremail    string `db:"useremail"`
}
type UserResp struct {
	Id           int    `json:"Id"`
	Username     string `json:"Username"`
	Userpassword string `json:"Userpassword"`
	Useremail    string `json:"Useremail"`
}
type Userlogin struct {
	Username     string `json:"Username"`
	Userpassword string `json:"Userpassword"`
}

type FilestatsDB struct {
	Id               int       `db:"id"`
	Totallines       int       `db:"totallines"`
	Totalwords       int       `timerdb:"totalwords"`
	Totalspaces      int       `db:"totalspaces"`
	Totalvowels      int       `db:"totalvowels"`
	Totalpunctuation int       `db:"totalpunctuation"`
	Timestamp        time.Time `db:"timestamp"`
}
type Filestats struct {
	Id               int       `json:"id"`
	Totallines       int       `json:"totallines"`
	Totalwords       int       `json:"totalwords"`
	Totalspaces      int       `json:"totalspaces"`
	Totalvowels      int       `json:"totalvowels"`
	Totalpunctuation int       `json:"totalpunctuation"`
	Timestamp        time.Time `json:"timestamp"`
}

// type Filestatistics struct {
// 	Linecount        int
// 	Wordscount       int
// 	Vowelscount      int
// 	Punctuationcount int
// }
