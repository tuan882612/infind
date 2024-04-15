package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNotFound      = errors.New("user not found")
	ErrAlreadyExists = errors.New("user already exists")
)

type User struct {
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	Email    string    `json:"email" validate:"required,email"`
	Password []byte    `json:"password" validate:"required"`
	Created  time.Time `json:"created" validate:"required"`
}

func NewUser(email, password string) (*User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	return &User{
		UserID:   uuid.New(),
		Email:    email,
		Password: []byte(hashed),
		Created:  time.Now(),
	}, nil
}
