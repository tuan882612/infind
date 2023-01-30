package response

import (
	"net/http"
)

type Base struct {
	Body map[string]string `json:"body"`
	Meta Meta              `json:"meta"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func None(messagge string) Base {
	return Base{
		Body: map[string]string{},
		Meta: Meta{
			Code: http.StatusNotFound,
			Message: messagge,
		},
	}
}

func Ok(messagge string, body map[string]string) Base {
	return Base{
		Body: body,
		Meta: Meta{
			Code: http.StatusOK,
			Message: messagge,
		},
	}
}
