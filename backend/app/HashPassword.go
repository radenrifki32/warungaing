package app

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func OpenPassword(hashing []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashing, password)
	return err
}
