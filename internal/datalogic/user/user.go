package user

import (
	"context"
	userModel "github.com/AndreeJait/template-service-go/internal/model/user"
	"time"
)

type DataLogic interface {
	AuthenticateUser(ctx context.Context, param userModel.AuthenticateUserParam) (user userModel.User, err error)
	GetTokenUser(ctx context.Context, user userModel.User) (token string, expiredAt time.Time, refreshToken string, err error)
}
