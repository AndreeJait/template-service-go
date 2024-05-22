package exam_browser

import (
	"context"
	"github.com/AndreeJait/exam-platform-service/internal/model/user"
)

type ExamBrowser interface {
	// users
	GetUserByID(ctx context.Context, userID string) (user user.User, err error)
}
