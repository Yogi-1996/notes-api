package hash

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyPassword(password, hashpassword string) bool {
	if password == "" || hashpassword == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	return err == nil
}
