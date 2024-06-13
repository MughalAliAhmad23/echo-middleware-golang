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
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"

	"github.com/labstack/echo"
)

func Add(c echo.Context) error {

	var input models.CalculatorReq

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := input.No1 + input.No2
	res, err := db.Insert(models.CalculatorDb{No1: input.No1,
		No2:       input.No2,
		Operation: "+",
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
		Message: "Successfully calculted addition",
		Status:  http.StatusOK,
	}
	return c.JSON(http.StatusCreated, resp)

}

func TextfilePro(c echo.Context) error {

	goVal := c.FormValue("goroutines")
	if goVal == "" {
		return c.JSON(http.StatusBadRequest, "value of goroutines is empty!")
	}
	goIntVal, err := strconv.Atoi(goVal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(goVal)
	var wg sync.WaitGroup

	defer filereader.Timer("main")()

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	defer src.Close()

	osfile := src.(*os.File)

	filedata, err := osfile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	filesize := filedata.Size()

	chunksize := filesize / int64(goIntVal)
	fmt.Println("chunk size", chunksize)

	reader := bufio.NewReader(osfile)

	chanResult := make(chan models.Filestats, goIntVal)

	for i := 0; i < goIntVal; i++ {
		fmt.Println("routine is ", runtime.NumCPU())
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

	for i := 0; i < goIntVal; i++ {
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
		return c.JSON(http.StatusInternalServerError, err.Error())
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
	return c.JSON(http.StatusCreated, resp)

}

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

func Division(c echo.Context) error {
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
