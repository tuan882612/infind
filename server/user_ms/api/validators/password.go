package validators

import "golang.org/x/crypto/bcrypt"

func ValidateHash(password string, hash string) bool {
	bPassword := []byte(password)
	bHash := []byte(hash)
	return bcrypt.CompareHashAndPassword(bPassword, bHash) == nil
}