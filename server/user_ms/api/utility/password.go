package utility

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytes, 14)
	return string(hash), err
}