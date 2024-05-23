package auth

import (
	"context"
	"github.com/AndreeJait/template-service-go/internal/model/auth"
)

type UseCase interface {
	Login(ctx context.Context, param auth.LoginParam) (resp auth.AuthenticatedResponse, err error)
}
