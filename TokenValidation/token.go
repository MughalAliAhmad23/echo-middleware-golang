package tokenvalidation

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func Isvalid(tokenstring string) error {

	if tokenstring == "" {
		return fmt.Errorf("missing Authorization Header")
	}

	newToken, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte("Secret-key"), nil
	})

	if err != nil {
		return fmt.Errorf("invalid header")
	}

	if !newToken.Valid {
		return fmt.Errorf("Un-Authorized")
	} else {
		fmt.Println("Authorized")
	}
	return nil
}
