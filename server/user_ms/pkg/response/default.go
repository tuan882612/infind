package response

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
