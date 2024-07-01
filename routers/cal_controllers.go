package routers

import (
	"bufio"
	"calculator/db"
	"calculator/filereader"
	"calculator/models"
	"calculator/myjwt"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]string)

//var broadcast = make(chan models.Message)

// Add godoc
// @Summary Add two numbers
// @Description Adds two numbers and stores the result in the database
// @Tags calculator
// @Accept  json
// @Produce  json
// @Param   input  body models.CalculatorReq  true  "Input data: No1 and No2 should nt be zero"
// @security BearerAuth
// @Success 200 {object} models.Resp "Successfully calculated addition"
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Router /calculator/add [post]
func Add(c echo.Context) error {

	// request:= c.Request()
	// responseWriter := c.

	var input models.CalculatorReq

	err := c.Bind(&input)
	if err != nil {
		fmt.Println("i am in bind section")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := input.No1 + input.No2
	res, err := db.Insert(models.CalculatorDb{No1: input.No1,
		No2:       input.No2,
		Operation: "+",
		Result:    result,
	})
	if err != nil {
		fmt.Println("i am in insert section")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	output := models.CalculatorResp{
		Id:        res.Id,
		No1:       res.No1,
		No2:       res.No2,
		Operation: res.Operation,
		Result:    res.Result,
	}
	resp := models.Resp{
		Data:    output,
		Message: "Successfully calculted addition",
		Status:  http.StatusOK,
	}
	fmt.Println("i am about to leave the add function")
	return c.JSON(http.StatusCreated, resp)

}

func Handlecnnections(c echo.Context) error {

	req := c.Request()
	headers := req.Header

	apitoken := headers.Get("Authorization")

	claims, _ := ExtractClaims(apitoken)
	fmt.Println(claims)
	name := fmt.Sprintf("%s", claims["username"])
	fmt.Println(name)

	conn, err := upgrade.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Println("connection err:", err)
		return err
	}
	defer conn.Close()

	clients[conn] = name
	// we know ws send continously data and recive so for that we make a for loop...for continously reading and writing
	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("error in reading message", err)
			delete(clients, conn)
			return err
		}
		//broadcast <- msg
	}

}
func Processesfile(gorountines int, file multipart.FileHeader, name string) {
	//valueExists := false
	var wg sync.WaitGroup
	src, err := file.Open()
	if err != nil {
		for client, username := range clients {
			if username == name {
				err := client.WriteJSON(err.Error())
				if err != nil {
					fmt.Println("err in client message", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
	defer src.Close()

	osfile := src.(*os.File)

	filedata, err := osfile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	filesize := filedata.Size()

	chunksize := filesize / int64(gorountines)
	fmt.Println("chunk size", chunksize)

	reader := bufio.NewReader(osfile)

	chanResult := make(chan models.Filestats, gorountines)

	for i := 0; i < gorountines; i++ {
		//fmt.Println("routine is ", goIntVal)
		chunk := make([]byte, chunksize)
		_, err := reader.Read(chunk)
		if err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		go filereader.FileProcessor(chunk, chanResult, &wg)
	}
	wg.Wait()

	totalResut := models.FilestatsDB{}

	for i := 0; i < gorountines; i++ {
		result := <-chanResult
		totalResut.Totallines += result.Totallines
		totalResut.Totalwords += result.Totalwords
		totalResut.Totalvowels += result.Totalvowels
		totalResut.Totalpunctuation += result.Totalpunctuation
	}

	res, err := db.Fileinsert(models.FilestatsDB{
		Totallines:       totalResut.Totallines,
		Totalwords:       totalResut.Totalwords,
		Totalspaces:      totalResut.Totalwords - 1,
		Totalvowels:      totalResut.Totalvowels,
		Totalpunctuation: totalResut.Totalpunctuation,
	})
	if err != nil {
		for client, username := range clients {
			if username == name {
				err := client.WriteJSON(err.Error())
				if err != nil {
					fmt.Println("err in client message", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
		//return c.JSON(http.StatusInternalServerError, err.Error())
	}

	output := models.Filestats{
		Id:               res.Id,
		Totallines:       res.Totallines,
		Totalwords:       res.Totalwords,
		Totalspaces:      res.Totalspaces,
		Totalvowels:      res.Totalvowels,
		Totalpunctuation: res.Totalpunctuation,
		Timestamp:        res.Timestamp,
	}

	resp := models.Resp{
		Data:    output,
		Message: "File result Successfully Added",
		Status:  http.StatusOK,
	}

	fmt.Println("Id :", res.Id)
	fmt.Println("Total lines :", res.Totallines)
	fmt.Println("Total words :", res.Totalwords)
	fmt.Println("Total spaces :", res.Totalspaces)
	fmt.Println("Total vowels :", res.Totalvowels)
	fmt.Println("Total Punctuation :", res.Totalpunctuation)
	fmt.Println("Timestamp of a file :", res.Timestamp)

	fmt.Println("main exists")

	for client, username := range clients {
		if username == name {
			err := client.WriteJSON(resp)
			if err != nil {
				fmt.Println("err in client message", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
	//return c.JSON(http.StatusCreated, resp)
}

// Textfilepro godooc
// @Summary Count the textfile stats
// @Description Get the text filr from the user and count the filestats like lines,spaces,words,vowels,punctuation & timestamps
// @Tags Textfile
// @Accept multipart/form-data
// @Produce json
// @Param goroutines formData string true "Number of goroutines"
// @Param file formData file true "Text file to process"
// @security BearerAuth
// @Success 200 {object} models.Resp "File result Successfully Added"
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Router /textfileprocessor [post]
func TextfilePro(c echo.Context) error {

	req := c.Request()
	headers := req.Header

	apitoken := headers.Get("Authorization")

	claims, _ := ExtractClaims(apitoken)
	fmt.Println(claims)
	name := fmt.Sprintf("%s", claims["username"])
	fmt.Println(name)

	goVal := c.FormValue("goroutines")
	if goVal == "" {
		return c.JSON(http.StatusBadRequest, "value of goroutines is empty!")
	}
	goIntVal, err := strconv.Atoi(goVal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(goVal)
	//var wg sync.WaitGroup

	defer filereader.Timer("main")()

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	go Processesfile(goIntVal, *file, name)
	resp := models.Resp{
		Data:    nil,
		Message: "File successfully uploaded!",
		Status:  http.StatusOK,
	}
	return c.JSON(http.StatusCreated, resp)

}

// Getallstats godoc
// @Summary Get all stats of a processed file
// @Description Get all the stats of a file stored in database like timestamp,lines,words,spaces,vowels
// @tags Textfile
// @Produce json
// @security BearerAuth
// @success 200 {array} models.Resp "Successfully Retrieved all data"
// @Failure 400 {object} string "Invalid input"
// @Router /textfilestats/all [Get]
func Getallstats(c echo.Context) error {
	filestats, err := db.Readallfilestats()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res := models.Resp{
		Data:    filestats,
		Message: "Successfully retrieves the result",
		Status:  http.StatusOK,
	}
	return c.JSON(http.StatusCreated, res)

}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := myjwt.SecretKey
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

// Crediantials godoc
// @Summary User Signup
// @Description Signup the user and stores the stats of a user in a database
// @Tags user
// @Accept json
// @Produce json
// @Param input body models.User true "username,userpassword & useremail cannot be empty"
// @Success 200 {object} models.Resp "User verified successfully"
// @Failure 400 {object} models.Resp "User already exists"
// @Failure 500 {object} string "Invalid input"
// @Router /user [Post]
func Crediantials(c echo.Context) error {
	var input models.User

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	exist, err := db.Isuserexists(models.User{
		Username: input.Username,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if exist {
		resp := models.Resp{
			Data:    nil,
			Message: "User already exists",
			Status:  http.StatusOK,
		}
		return c.JSON(http.StatusOK, resp)
	}

	res, err := db.UserInsert(models.UserDb{
		Username:     input.Username,
		Userpassword: input.Userpassword,
		Useremail:    input.Useremail,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	output := models.UserResp{
		Id:           res.Id,
		Username:     res.Username,
		Userpassword: res.Userpassword,
		Useremail:    res.Useremail,
	}

	resp := models.Resp{
		Data:    output,
		Message: "User verified successfully",
		Status:  http.StatusOK,
	}

	return c.JSON(http.StatusCreated, resp)
}

// Login godoc
// @Summary User Login
// @Description Login user and generates the Authorixation token
// @Tags user
// @Accept json
// @Produce json
// @Param input body models.Userlogin true "username & userpassword cannot be empty"
// @Success 200 {object} models.Resp "completed"
// @Failure 400 {object} models.Resp " "User does not exists""
// @Failure 500 {object} string "Invalid input"
// @Router /user/login [Post]
func Login(c echo.Context) error {
	var input models.Userlogin

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	exist, err := db.Userlogin(models.Userlogin{
		Username:     input.Username,
		Userpassword: input.Userpassword,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if !exist {
		resp := models.Resp{
			Data:    nil,
			Message: "user does not exists",
			Status:  http.StatusOK,
		}
		return c.JSON(http.StatusOK, resp)
	}

	token, err := myjwt.GenerateToken(input.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	resp := models.Resp{
		Data:    token,
		Message: "completed",
		Status:  http.StatusOK,
	}

	return c.JSON(http.StatusCreated, resp)
}

// Substract godoc
// @Summary Substract two numbers
// @Description Substract two numbers and stores the result in the database
// @Tags calculator
// @Accept  json
// @Produce  json
// @Param   input  body     models.CalculatorReq  true  "Input data: No1 and No2 should not be zero"
// @security BearerAuth
// @Success 200 {object} models.Resp "Successfully calculated substraction"
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Router /calculator/substract [post]
func Substract(c echo.Context) error {
	var input models.CalculatorReq

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := input.No1 - input.No2

	res, err := db.Insert(models.CalculatorDb{No1: input.No1,
		No2:       input.No2,
		Operation: "-",
		Result:    result,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	output := models.CalculatorResp{
		Id:        res.Id,
		No1:       res.No1,
		No2:       res.No2,
		Operation: res.Operation,
		Result:    res.Result,
	}

	resp := models.Resp{
		Data:    output,
		Message: "Successfully calculted substraction",
		Status:  http.StatusOK,
	}

	return c.JSON(http.StatusCreated, resp)
}

// Multiply godoc
// @Summary Multiply two numbers
// @Description Multiplys two numbers and stores the result in the database
// @Tags calculator
// @Accept  json
// @Produce  json
// @Param   input  body models.CalculatorReq  true  "Input data: No1 and No2 should nt be zero"
// @security BearerAuth
// @Success 200 {object} models.Resp "Successfully calculated multiplication"
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Router /calculator/multiply [post]
func Multiply(c echo.Context) error {

	var input models.CalculatorReq

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := input.No1 * input.No2
	res, err := db.Insert(models.CalculatorDb{No1: input.No1,
		No2:       input.No2,
		Operation: "*",
		Result:    result,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	output := models.CalculatorResp{
		Id:        res.Id,
		No1:       res.No1,
		No2:       res.No2,
		Operation: res.Operation,
		Result:    res.Result,
	}

	resp := models.Resp{
		Data:    output,
		Message: "Successfully calculated multiplication",
		Status:  http.StatusOK,
	}

	return c.JSON(http.StatusCreated, resp)
}

// Divide godoc
// @Summary Divide two numbers
// @Description Divides two numbers and stores the result in the database
// @Tags calculator
// @Accept  json
// @Produce  json
// @Param   input  body  models.CalculatorReq  true  "Input data: No1 and No2 should not be zero"
// @security BearerAuth
// @Success 200 {object} models.Resp "Successfully calculated division"
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Router /calculator/divide [post]
func Divide(c echo.Context) error {
	var input models.CalculatorReq

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := input.No1 / input.No2

	res, err := db.Insert(models.CalculatorDb{No1: input.No1,
		No2:       input.No2,
		Operation: "/",
		Result:    result,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	output := models.CalculatorResp{
		Id:        res.Id,
		No1:       res.No1,
		No2:       res.No2,
		Operation: res.Operation,
		Result:    res.Result,
	}

	resp := models.Resp{
		Data:    output,
		Message: "Successfully calculated division",
		Status:  http.StatusOK,
	}

	return c.JSON(http.StatusCreated, resp)
}

// Getall  godoc
// @Summary Retrives all record
// @Description Retrives all records from calculator table in the database
// @Tags calculator
// @Produce json
// @security BearerAuth
// @success 200 {array} models.Resp "Successfully Retrieved all data"
// @Failure 400 {object} string "Invalid input"
// @Router /calculator [Get]
func Getall(c echo.Context) error {
	calculations, err := db.Readall()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp := models.Resp{
		Data:    calculations,
		Message: "Successfully Retrieved all data",
		Status:  http.StatusOK,
	}

	return c.JSON(http.StatusCreated, resp)
}

// Getbyid godoc
// @Summary Retrives recored on specific id
// @Description Retrives record on specific id from calculator table in the database
// @Tags calculator
// @Produce json
// @Param id path string true "id cannot be empty"
// @security BearerAuth
// @Success 200 {object} models.Resp "Successfully retrieved id data"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Internal server error"
// @Router /calculator/{id} [Get]
func Getbyid(c echo.Context) error {
	request_id := c.Param("id")

	if request_id == "" {
		return c.JSON(http.StatusBadRequest, "id can't be empty")
	}

	calculationID, err := strconv.Atoi(request_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cal, err := db.Readbyid(calculationID)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "invalid ID")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	resp := models.Resp{
		Data:    cal,
		Message: "Successfully retrieved id data",
		Status:  http.StatusOK,
	}

	return c.JSON(http.StatusCreated, resp)
}

// Getbysymbol godoc
// @Summary Retrives all recored on specific operation
// @Description Retrives all record on specific operation from calculator table in the database
// @Tags calculator
// @Produce json
// @Param operation path string true "operation can be in decodes foam"
// @security BearerAuth
// @Success 200 {object} models.Resp "Successfully retrieved symbol data"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Internal server error"
// @Router /calculator/symbol{operation} [Get]
func Getbysymbol(c echo.Context) error {

	request_symbol := c.QueryParam("operation")
	if request_symbol == "" {
		return c.JSON(http.StatusBadRequest, "symbol can't be empty")
	}

	fmt.Println(request_symbol)

	res, err := db.Readbysymbol(request_symbol)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	resp := models.Resp{
		Data:    res,
		Message: "retrieved all symbol data",
		Status:  http.StatusOK,
	}

	return c.JSON(http.StatusCreated, resp)
}

// Delete godoc
// @Summary Delete the recored on specific id
// @Description Delete the record on specific id from calculator table in the database
// @Tags calculator
// @Produce json
// @Param id path string true "id cannot be empty"
// @security BearerAuth
// @Success 200 {object} models.Resp "Successfully deleted"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Internal server error"
// @Router /calculator/{id} [Delete]
func Delete(c echo.Context) error {
	request_id := c.Param("id")

	if request_id == "" {
		return c.JSON(http.StatusBadRequest, "id can't be empty")
	}

	calculationID, err := strconv.Atoi(request_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = db.Removebyid(calculationID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	resp := models.Resp{
		Data:    nil,
		Message: "Successfully Deleted",
		Status:  http.StatusOK,
	}

	return c.JSON(http.StatusCreated, resp)
}

// Update godoc
// @Summary Update the recored on specific id
// @Description Update the record on specific id from calculator table in the database
// @Tags calculator
// @Accept json
// @Produce json
// @Param input body models.CalculatorReq true "Input data: No2 cannot be zero"
// @Param id path string true "id cannot be empty"
// @security BearerAuth
// @Success 200 {object} models.Resp "Successfully deleted"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Internal server error"
// @Router /calculator/{id} [Put]
func Update(c echo.Context) error {
	var cal models.CalculatorReq

	err := c.Bind(&cal)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	request_id := c.Param("id")
	if request_id == "" {
		return c.JSON(http.StatusBadRequest, "id can't be empty")
	}

	calculation_ID, err := strconv.Atoi(request_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	existingsymbol, err := db.Readbyid(calculation_ID)
	if err != nil {
		return err
	}

	var result float64

	switch existingsymbol.Operation {
	case "+":
		result = cal.No1 + cal.No2
	case "-":
		result = cal.No1 - cal.No2
	case "*":
		result = cal.No1 * cal.No2
	case "/":
		if cal.No2 == 0 {
			return errors.New("division by zero")
		}
		result = cal.No1 / cal.No2
	default:
		return errors.New("unsupported operation")
	}

	fmt.Println(result)

	err = db.Reupdate(calculation_ID, float64(cal.No1), float64(cal.No2), float64(result))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	output := models.CalculatorResp{
		Id:        calculation_ID,
		No1:       cal.No1,
		No2:       cal.No2,
		Operation: existingsymbol.Operation,
		Result:    result,
	}

	resp := models.Resp{
		Data:    output,
		Message: "Updated Successfully",
		Status:  http.StatusOK,
	}

	return c.JSON(http.StatusCreated, resp)
}
