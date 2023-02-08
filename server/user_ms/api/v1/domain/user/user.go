package user

import "userms/api/v1/domain/search"

type User struct {
	Username string          `dynamodbav:"username" json:"username"`
	Email    string          `dynamodbav:"email" json:"email"`
	Password string          `dynamodbav:"password" json:"password"`
	Created  string          `dynamodbav:"created" json:"created"`
	History  []search.Search `dynamodbav:"history" json:"history"`
}
