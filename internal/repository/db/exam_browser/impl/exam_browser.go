package impl

import (
	"github.com/AndreeJait/template-service-go/internal/repository/db/exam_browser"
	"github.com/AndreeJait/go-utility/loggerw"
	"github.com/AndreeJait/go-utility/sqlw/postgres"
)

type examBrowser struct {
	sqlW postgres.SqlW
	log  loggerw.Logger
}

func NewRepository(sqlW postgres.SqlW, log loggerw.Logger) exam_browser.ExamBrowser {
	return &examBrowser{
		sqlW: sqlW,
		log:  log,
	}
}
