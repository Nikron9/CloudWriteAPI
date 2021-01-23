package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func CipherPassword(password string) (string, error) {
	cipheredPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(cipheredPassword), err
}

func CheckPassword(password string, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
