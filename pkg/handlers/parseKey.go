package handlers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func ParseKey(raw []byte, method string) (any, error) {
	switch {
	case strings.HasPrefix(method, "rs"):
		return jwt.ParseRSAPrivateKeyFromPEM(raw)
	case strings.HasPrefix(method, "hs"):
		return raw, nil
	case strings.HasPrefix(method, "es"):
		return jwt.ParseECPrivateKeyFromPEM(raw)
	case strings.HasPrefix(method, "eddsa"):
		return jwt.ParseEdPrivateKeyFromPEM(raw)
	}
	return nil, fmt.Errorf("method %s not supported", method)
}
