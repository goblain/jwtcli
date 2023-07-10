package handlers

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt/v5"
	"strings"
)

func GetSigningMethod(label string) (jwt.SigningMethod, error) {
	var err error
	method := jwt.GetSigningMethod(strings.ToUpper(label))
	if method == nil {
		err = fmt.Errorf("%s did not return a valid signing method", label)
	}
	return method, err
}
