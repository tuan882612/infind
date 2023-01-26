package model

type UserResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Body    User   `json:"body"`
}

type DefaultResponse struct {
	Code    int               `json:"code"`
	Message string 		      `json:"message"`
	Body    map[string]string `json:"body"`
}

type LoginResponse struct {
	Found    bool   `json:"found"`
}