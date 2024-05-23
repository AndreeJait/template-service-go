package impl

import (
	"context"
	"github.com/AndreeJait/go-utility/jwt"
	userModel "github.com/AndreeJait/template-service-go/internal/model/user"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"time"
)

func (d *dataLogic) GetTokenUser(ctx context.Context, user userModel.User) (token string, expiredAt time.Time, refreshToken string, err error) {
	expiredAt = d.timeInternal.Now().Add(time.Duration(d.cfg.Setting.JwtTokenDurationExpiredInMinute) * time.Minute)
	claims := userModel.InternalClaims{
		Claims: jwt2.RegisteredClaims{
			Issuer:    d.cfg.Service.Name,
			ExpiresAt: jwt2.NewNumericDate(expiredAt),
		},
		Email:  user.GetUsername(),
		Role:   user.GetRole(),
		UserID: user.GetUserID(),
	}
	token, err = jwt.CreateToken(jwt.CreateTokenRequest{
		Claims:      claims,
		SecretToken: d.cfg.Secret.Jwt,
	})
	if err != nil {
		return
	}

	timeDurationRefresh := d.timeInternal.Now().Add(time.Duration(d.cfg.Setting.JwtRefreshTokenDurationExpiredInMinute) * time.Minute)
	claims = userModel.InternalClaims{
		Claims: jwt2.RegisteredClaims{
			Issuer:    d.cfg.Service.Name,
			ExpiresAt: jwt2.NewNumericDate(timeDurationRefresh),
		},
		UserID: user.GetUserID(),
	}
	refreshToken, err = jwt.CreateToken(jwt.CreateTokenRequest{
		Claims:      claims,
		SecretToken: d.cfg.Secret.Jwt,
	})
	if err != nil {
		return
	}
	return
}
