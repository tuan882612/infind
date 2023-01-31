package response

import (
	"net/http"
	"userms/api/v1/model"
)

type User struct {
	Body model.User `json:"body"`
	Meta Meta		`json:"meta"`
}

func FoundUser(user model.User) User {
	return User{
		Body: user,
		Meta: Meta{
			Code: http.StatusOK,
			Message: "",
		},
	}
}