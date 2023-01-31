package response

import (
	"net/http"
	"userms/api/v1/model"
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

func Error(err error) Base {
	return Base{
		Body: map[string]string{},
		Meta: Meta{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		},
	}
}

func Created(user model.User) Base {
	return Base{
		Body: map[string]string{},
		Meta: Meta{
			Code: http.StatusCreated,
			Message: "",
		},
	}
}

func Updated(user model.User) Base {
	return Base{
		Body: map[string]string{},
		Meta: Meta{
			Code: http.StatusAccepted,
			Message: "",
		},
	}
}

func Exist(user model.User) Base {
	return Base{
		Body: map[string]string{},
		Meta: Meta{
			Code: http.StatusConflict,
			Message: "",
		},
	}
}

func Accepted(msg string) Base {
	return Base{
		Body: map[string]string{},
		Meta: Meta{
			Code: http.StatusAccepted,
			Message: msg,
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
