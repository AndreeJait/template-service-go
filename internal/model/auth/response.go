package auth

type AuthenticatedResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredAt    string `json:"expired_at"`
	Role         string `json:"role"`
}
