package tokenvalidation

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func Isvalid(tokenstring string) error {
	if tokenstring == "" {
		fmt.Println("Missing Authorization Header")
		return fmt.Errorf("missing Authorization Header")
	}
	newToken, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte("Secret-key"), nil
	})

	if err != nil {
		fmt.Println("im in invalid err section")
		return fmt.Errorf("invalid header")
	}

	if !newToken.Valid {
		fmt.Println("im in invalid not valid section")
		fmt.Println("Un-Authorized")
		return fmt.Errorf("Un-Authorized")
	} else {
		fmt.Println("im in isvalid esle section")
		fmt.Println("Authorized")
	}
	return nil
}
