package impl

import (
	"context"
	"github.com/AndreeJait/template-service-go/internal/model/user"
)

func (e *testingImpl) GetUserByID(ctx context.Context, userID string) (user user.User, err error) {
	err = e.sqlW.Get(ctx,
		&user,
		queryGetUserByID,
		userID)
	return
}
