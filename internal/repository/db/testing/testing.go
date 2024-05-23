package testing

import (
	"context"
	"github.com/AndreeJait/template-service-go/internal/model/user"
)

type Testing interface {
	// GetUserByID get user by user_id
	GetUserByID(ctx context.Context, userID string) (user user.User, err error)

	// GetUserByEmail get user by email
	GetUserByEmail(ctx context.Context, email string) (user user.User, err error)
}
