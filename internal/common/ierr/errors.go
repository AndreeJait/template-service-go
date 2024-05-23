package ierr

import "github.com/AndreeJait/go-utility/errow"

const (
	ErrCodeUserNotFound errow.ErrorWCode = iota + 400001
)

const (
	ErrCodePasswordNotMatch errow.ErrorWCode = iota + 401001
)

var (
	ErrUserNotFound = errow.ErrorW{Code: ErrCodeUserNotFound, Message: "user not found"}

	ErrPasswordNotMatch = errow.ErrorW{Code: ErrCodePasswordNotMatch, Message: "password not match"}
)
