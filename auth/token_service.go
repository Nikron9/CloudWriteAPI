package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user string, email string) (string, error) {
	exp, err := time.ParseDuration(tokenConfig.tokenExp)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user,
		"email":    email,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(exp).Unix(),
	})
	s, err := token.SignedString([]byte(tokenConfig.tokenSecret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func verifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("There was an error")
		}
		return []byte(tokenConfig.tokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid == true {
		return token.Claims.(jwt.MapClaims), err
	} else {
		return nil, fmt.Errorf("Token invalid.")
	}
}
