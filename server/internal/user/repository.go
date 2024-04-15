package user

import (
	"context"
	"errors"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type Repository struct {
	postgres *pgxpool.Pool
	userPool sync.Pool
}

func NewRepository(postgres *pgxpool.Pool) *Repository {
	return &Repository{
		postgres: postgres,
		userPool: sync.Pool{
			New: func() interface{} {
				return &User{}
			},
		},
	}
}

func (r *Repository) CreateTx(ctx context.Context) (pgx.Tx, error) {
	tx, err := r.postgres.Begin(ctx)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	return tx, nil
}

func (r *Repository) GetUser(ctx context.Context, email string) (*User, error) {
	user := r.userPool.Get().(*User)
	defer r.userPool.Put(user)

	row := r.postgres.QueryRow(ctx, GetUserQuery, email)
	if err := row.Scan(&user.UserID, &user.Email, &user.Password); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		log.Error().Err(err).Send()
		return nil, err
	}

	return user, nil
}

func (r *Repository) Register(ctx context.Context, tx pgx.Tx, user *User) error {
	_, err := tx.Exec(ctx, RegisterQuery, &user.UserID, &user.Email, &user.Password, &user.Created)
	if err != nil {
		pgErr := &pgconn.PgError{}
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrAlreadyExists
		}

		log.Error().Err(err).Send()
		return err
	}

	return nil
}
