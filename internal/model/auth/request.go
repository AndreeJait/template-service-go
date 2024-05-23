package auth

type LoginParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
