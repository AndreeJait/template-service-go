package impl

import (
	"context"
	"github.com/AndreeJait/template-service-go/common/constant"
	"github.com/AndreeJait/template-service-go/internal/model/user"
)

func (e *testingImpl) GetUserByEmail(ctx context.Context, email string) (user user.User, err error) {
	if e.appMode == constant.Development {
		e.log.Infof("execute sql %s", queryGetUserByEmail)
	}
	err = e.sqlW.Get(ctx,
		&user,
		queryGetUserByEmail,
		email)
	return
}
