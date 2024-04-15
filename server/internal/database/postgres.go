package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func NewPostgres(pgUrl string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) 
	defer cancel()

	config, err := pgxpool.ParseConfig(pgUrl)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	log.Info().Msg("postgres connected...")
	return pool, nil
}
