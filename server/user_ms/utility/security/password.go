package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytes, 14)
	return string(hash), err
}

func ValidateHash(password string, hash *string) bool {
	bPassword := []byte(password)
	bHash := []byte(*hash)
	return bcrypt.CompareHashAndPassword(bHash, bPassword) == nil
}