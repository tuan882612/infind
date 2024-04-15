package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"infind/internal/auth"
)

type Server struct {
	router *echo.Echo
	addr   string
	deps   *Depends
	errCh  chan error
}

func New(deps *Depends) (*Server, error) {
	svr := &Server{
		router: echo.New(),
		addr:   deps.Config.Server.Address,
		deps:   deps,
		errCh:  make(chan error),
	}
	svr.setupMiddlewares()	
	svr.setupRoutes()
	return svr, nil
}

func (s *Server) setupMiddlewares() {
	s.router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod: true, LogURI: true, LogStatus: true, LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmtstring := "request: %s %s, status: %d, latency: %s"
			log.Info().Msg(fmt.Sprintf(fmtstring, v.Method, v.URI, v.Status, v.Latency))
			return nil
		},
	}))
	s.router.Use(middleware.CORS())
	s.router.Use(middleware.RequestID())
}

func (s *Server) setupRoutes() {
	s.router.GET("/api/health", func(c echo.Context) error {
		body := map[string]string{"status": "ok"}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		return jsoniter.NewEncoder(c.Response().Writer).Encode(body)
	})

	authHandler := auth.NewHandler(s.deps.AuthSvc, s.deps.JWTProvider)
	authRoutes := s.router.Group("/api/auth")
	authRoutes.POST("/login", authHandler.Login)
	authRoutes.POST("/register", authHandler.Register)
}

func (s *Server) Start() error {
	go s.watch()

	if err := s.router.Start(s.addr); err != http.ErrServerClosed {
		log.Error().Err(err).Send()
		return err
	}

	if err, ok := <-s.errCh; ok {
		log.Error().Err(err).Send()
		return err
	}

	return nil
}

func (s *Server) watch() {
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
	<-sigch

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.deps.Postgres.Close()

	s.errCh <- s.router.Shutdown(ctx)
	close(s.errCh)
}
