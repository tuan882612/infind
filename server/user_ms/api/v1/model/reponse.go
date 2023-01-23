package model

type Response struct {
	Code    int 			  `json:"code"`
	Message string 			  `json:"message"`
	Body    map[string]string `json:"body"`
}