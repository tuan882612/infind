package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"infind/internal/server"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}).With().Caller().Logger()

	deps, err := server.NewDepends()
	if err != nil {
		log.Fatal().Msg("failed to create dependencies")
	}

	svr, err := server.New(deps)
	if err != nil {
		log.Fatal().Msg("failed to create server")
	}

	if err := svr.Start(); err != nil {
		log.Fatal().Msg("failed to start server")
	}

	log.Info().Msg("server shutdown gracefully...")
}
