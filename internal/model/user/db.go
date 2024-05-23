package user

type User struct {
	UserID   int64  `json:"user_id" db:"user_id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	FullName string `json:"full_name" db:"full_name"`
	Password string `json:"-" db:"password"`
	Role     string `json:"role" db:"role"`
}

func (u User) GetUserID() int64 {
	return u.UserID
}
func (u User) GetUsername() string {
	return u.Email
}
func (u User) GetRole() string {
	return u.Role
}
