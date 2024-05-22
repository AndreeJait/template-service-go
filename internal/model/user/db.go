package user

type User struct {
	UserID   int64  `json:"user_id" db:"user_id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Role     string `json:"role" db:"role"`
}
