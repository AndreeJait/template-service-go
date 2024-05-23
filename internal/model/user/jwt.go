package user

import "github.com/golang-jwt/jwt/v4"

type InternalClaims struct {
	jwt.Claims
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}
