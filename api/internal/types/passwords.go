package types

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	salt = 10
)

func IsPasswordValid(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func EncryptPassword(password string) (string, error) {
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(passwordByte), err
}
