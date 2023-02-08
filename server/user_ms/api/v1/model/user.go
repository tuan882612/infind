package model

type User struct {
	Username string   `dynamodbav:"username" json:"username"`
	Email    string   `dynamodbav:"email" json:"email"`
	Password string   `dynamodbav:"password" json:"password"`
	Created  string   `dynamodbav:"created" json:"created"`
	History  []Search `dynamodbav:"history" json:"history"`
}
