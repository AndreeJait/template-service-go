package impl

import (
	"github.com/AndreeJait/go-utility/loggerw"
	"github.com/AndreeJait/go-utility/timew"
	"github.com/AndreeJait/template-service-go/config"
	"github.com/AndreeJait/template-service-go/internal/datalogic/user"
	testingRepo "github.com/AndreeJait/template-service-go/internal/repository/db/testing"
)

type dataLogic struct {
	log          loggerw.Logger
	cfg          *config.Config
	testingRepo  testingRepo.Testing
	timeInternal timew.Time
}

func NewDataLogic(log loggerw.Logger, cfg *config.Config, testing testingRepo.Testing, timeInternal timew.Time) user.DataLogic {
	return &dataLogic{
		log:          log,
		cfg:          cfg,
		testingRepo:  testing,
		timeInternal: timeInternal,
	}
}
