package response

import "net/http"

func Error(err error) Base {
	return Base{
		Body: map[string]string{},
		Meta: Meta{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		},
	}
}