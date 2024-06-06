package routers

import (
	"calculator/db"
	"calculator/filereader"
	"calculator/models"
	"calculator/myjwt"
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo"
)

func Add(c echo.Context) error {
	// req := c.Request()
	// headers := req.Header

	// apitoken := headers.Get("Authorization")

	// err, verify := tokenvalidation.Isvalid(apitoken)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

	// if !verify {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

	// if apitoken == "" {
	// 	fmt.Println("Missing authorization header")
	// 	return c.JSON(http.StatusUnauthorized, "Missing authorization header")
	// }

	// newToken, err := jwt.Parse(apitoken, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte("Secret-key"), nil
	// })

	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, "Invalid header")
	// }

	// if !newToken.Valid {
	// 	fmt.Println("Un-Authorized")
	// 	return c.JSON(http.StatusUnauthorized, "Un-Authorized")
	// } else {
	// 	fmt.Println("Authorized")
	// }

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

	resp := make(chan string, 5)

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
	filedata, err := ioutil.ReadAll(src)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	wg.Add(1)
	go filereader.Wordfrequeny(string(filedata), &wg, resp)
	wg.Add(1)
	go filereader.SpaceCounter(string(filedata), &wg, resp)
	wg.Add(1)
	go filereader.Wordcounter(string(filedata), &wg, resp)
	wg.Add(1)
	go filereader.VowelsCounter(string(filedata), &wg, resp)
	wg.Add(1)
	go filereader.LineCounter(string(filedata), &wg, resp)

	wg.Wait()
	close(resp)
	for val := range resp {
		fmt.Println(val)
	}
	fmt.Println("iteratig iver channel")
	fmt.Println(len(resp))
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-resp)
	// }
	fmt.Println("main exists")
	return nil
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

	// req := c.Request()
	// headers := req.Header

	// apitoken := headers.Get("Authorization")

	// err, verify := tokenvalidation.Isvalid(apitoken)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

	// if !verify {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

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

	// req := c.Request()
	// headers := req.Header

	// apitoken := headers.Get("Authorization")

	// err, verify := tokenvalidation.Isvalid(apitoken)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

	// if !verify {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

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

	// req := c.Request()
	// headers := req.Header

	// apitoken := headers.Get("Authorization")

	// err, verify := tokenvalidation.Isvalid(apitoken)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

	// if !verify {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

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

	// req := c.Request()
	// headers := req.Header

	// apitoken := headers.Get("Authorization")

	// err, verify := tokenvalidation.Isvalid(apitoken)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

	// if !verify {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

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

	// req := c.Request()
	// headers := req.Header

	// apitoken := headers.Get("Authorization")

	// err, verify := tokenvalidation.Isvalid(apitoken)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

	// if !verify {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

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

	// req := c.Request()
	// headers := req.Header

	// apitoken := headers.Get("Authorization")

	// err, verify := tokenvalidation.Isvalid(apitoken)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

	// if !verify {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

	request_symbol := c.Param("operation")
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

	// req := c.Request()
	// headers := req.Header

	// apitoken := headers.Get("Authorization")

	// err, verify := tokenvalidation.Isvalid(apitoken)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

	// if !verify {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

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

	// req := c.Request()
	// headers := req.Header

	// apitoken := headers.Get("Authorization")

	// err, verify := tokenvalidation.Isvalid(apitoken)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

	// if !verify {
	// 	return c.JSON(http.StatusUnauthorized, "Missing Authorization Header!")
	// }

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
