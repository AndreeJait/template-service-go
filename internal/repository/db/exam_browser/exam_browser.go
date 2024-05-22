package exam_browser

import (
	"context"
	"github.com/AndreeJait/template-service-go/internal/model/user"
)

type ExamBrowser interface {
	// users
	GetUserByID(ctx context.Context, userID string) (user user.User, err error)
}
