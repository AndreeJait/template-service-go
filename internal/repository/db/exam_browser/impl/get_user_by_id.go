package impl

import (
	"context"
	"github.com/AndreeJait/exam-platform-service/internal/model/user"
)

func (e *examBrowser) GetUserByID(ctx context.Context, userID string) (user user.User, err error) {
	err = e.sqlW.Get(ctx,
		&user,
		queryGetUserByID,
		userID)
	return
}
