package impl

import (
	"github.com/AndreeJait/go-utility/loggerw"
	"github.com/AndreeJait/go-utility/sqlw/postgres"
	"github.com/AndreeJait/template-service-go/common/constant"
	"github.com/AndreeJait/template-service-go/config"
	"github.com/AndreeJait/template-service-go/internal/repository/db/testing"
)

type testingImpl struct {
	sqlW    postgres.SqlW
	log     loggerw.Logger
	cfg     *config.Config
	appMode constant.AppMode
}

func NewRepository(sqlW postgres.SqlW, log loggerw.Logger, cfg *config.Config, mode constant.AppMode) testing.Testing {
	return &testingImpl{
		sqlW:    sqlW,
		log:     log,
		appMode: mode,
		cfg:     cfg,
	}
}
