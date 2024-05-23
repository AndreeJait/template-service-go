package impl

import (
	"context"
	"github.com/AndreeJait/template-service-go/internal/model/auth"
	userModel "github.com/AndreeJait/template-service-go/internal/model/user"
	"time"
)

func (u *useCase) Login(ctx context.Context, param auth.LoginParam) (resp auth.AuthenticatedResponse, err error) {

	user, err := u.dlUser.AuthenticateUser(ctx, userModel.AuthenticateUserParam{
		Password: param.Password,
		Email:    param.Email,
	})
	if err != nil {
		return
	}

	token, expiredAt, refreshToken, err := u.dlUser.GetTokenUser(ctx, user)
	if err != nil {
		return
	}

	resp.Token = token
	resp.RefreshToken = refreshToken
	resp.ExpiredAt = expiredAt.Format(time.DateTime)
	resp.Role = user.GetRole()
	return resp, nil
}
