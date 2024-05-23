package impl

import (
	"context"
	"database/sql"
	"github.com/AndreeJait/go-utility/password"
	"github.com/AndreeJait/template-service-go/internal/common/ierr"
	userModel "github.com/AndreeJait/template-service-go/internal/model/user"
	"github.com/pkg/errors"
)

func (d *dataLogic) AuthenticateUser(ctx context.Context, param userModel.AuthenticateUserParam) (user userModel.User, err error) {
	user, err = d.testingRepo.GetUserByEmail(ctx, param.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, ierr.ErrUserNotFound
		}
		return
	}
	if !password.ComparePasswords(user.Password, []byte(param.Password)) {
		return user, ierr.ErrPasswordNotMatch
	}
	return user, nil
}
