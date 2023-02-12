package response

import (
	"net/http"
)

type Base struct {
	Body any  `json:"body"`
	Meta Meta `json:"meta"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Custom(body any, code int, message string) Base {
	return Base{
		Body: body,
		Meta: Meta{
			Code:    code,
			Message: message,
		},
	}
}

func Error(err error) Base {
	return Base{
		Body: map[string]string{},
		Meta: Meta{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		},
	}
}
