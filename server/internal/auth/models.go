package auth

import (
	"github.com/go-playground/validator"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type Request struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *Request) Process(e echo.Context) error {
	err := jsoniter.NewDecoder(e.Request().Body).Decode(r)
	if err != nil {
		return e.JSON(500, map[string]string{"error": err.Error()})
	}

	if err := validator.New().Struct(r); err != nil {
		return e.JSON(400, map[string]string{"error": err.Error()})
	}

	return nil
}
