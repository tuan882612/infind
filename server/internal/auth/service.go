package auth

import (
	"context"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"infind/internal/user"
)

type Service interface {
	Login(ctx context.Context, email, password string) (*user.User, error)
	Register(ctx context.Context, email, password string) (*user.User, error)
}

type service struct {
	repo *user.Repository
}

func NewService(repo *user.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Login(ctx context.Context, email, password string) (*user.User, error) {
	user, err := s.repo.GetUser(ctx, email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	return user, nil
}

func (s *service) Register(ctx context.Context, email, password string) (*user.User, error) {
	user, err := user.NewUser(email, password)
	if err != nil {
		return nil, err
	}

	tx, err := s.repo.CreateTx(ctx)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Register(ctx, tx, user); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	return user, nil
}
