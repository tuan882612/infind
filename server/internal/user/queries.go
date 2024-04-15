package user

const (
	GetUserQuery  = `SELECT user_id, email, password, created FROM users WHERE email = $1`
	RegisterQuery = `INSERT INTO users (user_id, email, password, created) VALUES ($1, $2, $3, $4)`
)
