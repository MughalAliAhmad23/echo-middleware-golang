package models

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
