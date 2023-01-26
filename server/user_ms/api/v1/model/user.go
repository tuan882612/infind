package model

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Created  string `json:"created"`
	History  []Search  `json:"history"`
}