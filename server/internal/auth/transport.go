package auth

import (
	"errors"
	"sync"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"infind/internal/jwt"
	"infind/internal/user"
)

type Handler struct {
	authSvc     Service
	jwtProvider *jwt.Provider
	reqPool     sync.Pool
}

func NewHandler(authSvc Service, jwtProvider *jwt.Provider) *Handler {
	return &Handler{
		authSvc:     authSvc,
		jwtProvider: jwtProvider,
		reqPool: sync.Pool{
			New: func() interface{} {
				return &Request{}
			},
		},
	}
}

func (h *Handler) Login(e echo.Context) error {
	req := h.reqPool.Get().(*Request)
	defer h.reqPool.Put(req)

	if err := req.Process(e); err != nil {
		return err
	}

	userInfo, err := h.authSvc.Login(e.Request().Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return e.JSON(404, map[string]string{"error": err.Error()})
		}

		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return e.JSON(401, map[string]string{"error": err.Error()})
		}

		return e.JSON(500, map[string]string{"error": err.Error()})
	}

	token, err := h.jwtProvider.GenerateToken(userInfo)
	if err != nil {
		return e.JSON(500, map[string]string{"error": err.Error()})
	}

	e.Response().Header().Set("Authorization", token)
	e.Response().Header().Set("X-User-ID", userInfo.UserID.String())
	e.Response().Header().Set("X-Email", userInfo.Email)
	return e.JSON(201, nil)
}

func (h *Handler) Register(e echo.Context) error {
	req := h.reqPool.Get().(*Request)
	defer h.reqPool.Put(req)

	if err := req.Process(e); err != nil {
		return err
	}

	userInfo, err := h.authSvc.Register(e.Request().Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, user.ErrAlreadyExists) {
			return e.JSON(409, map[string]string{"error": err.Error()})
		}

		return e.JSON(500, map[string]string{"error": err.Error()})
	}

	token, err := h.jwtProvider.GenerateToken(userInfo)
	if err != nil {
		return e.JSON(500, map[string]string{"error": err.Error()})
	}

	e.Response().Header().Set("Authorization", token)
	e.Response().Header().Set("X-User-ID", userInfo.UserID.String())
	e.Response().Header().Set("X-Email", userInfo.Email)
	return e.JSON(201, nil)
}
