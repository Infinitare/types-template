package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"types-template/cookies"
	"types-template/secrets"
)

func GetClaim(r *http.Request, key string) (interface{}, error) {

	auth := cookies.Get(r, "authorization")
	if auth == nil {
		return nil, errors.New("no authorization cookie")
	}

	token, err := jwt.Parse(auth.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(secrets.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("claims not mapclaims")
	}

	return claims[key], nil

}
