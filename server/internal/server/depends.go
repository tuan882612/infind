package server

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"infind/internal/auth"
	"infind/internal/config"
	"infind/internal/database"
	"infind/internal/jwt"
	"infind/internal/user"
)

type Depends struct {
	Config      *config.Configuration
	Postgres    *pgxpool.Pool
	JWTProvider *jwt.Provider
	AuthSvc     auth.Service
	UserRepo    *user.Repository
}

func NewDepends() (*Depends, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	postgres, err := database.NewPostgres(cfg.Database.PgURL)
	if err != nil {
		return nil, err
	}

	jwtProvider, err := jwt.NewProvider(cfg.JWT.Secret)
	if err != nil {
		return nil, err
	}

	userRepo := user.NewRepository(postgres)
	authSvc := auth.NewService(userRepo)

	return &Depends{
		Config:      cfg,
		Postgres:    postgres,
		JWTProvider: jwtProvider,
		AuthSvc:     authSvc,
		UserRepo:    userRepo,
	}, nil
}
