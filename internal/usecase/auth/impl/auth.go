package impl

import (
	"github.com/AndreeJait/go-utility/loggerw"
	"github.com/AndreeJait/go-utility/timew"
	"github.com/AndreeJait/template-service-go/config"
	dlUser "github.com/AndreeJait/template-service-go/internal/datalogic/user"
	testingRepo "github.com/AndreeJait/template-service-go/internal/repository/db/testing"
	"github.com/AndreeJait/template-service-go/internal/usecase/auth"
)

type useCase struct {
	log          loggerw.Logger
	cfg          *config.Config
	testingRepo  testingRepo.Testing
	timeInternal timew.Time
	dlUser       dlUser.DataLogic
}

func NewUseCase(log loggerw.Logger, cfg *config.Config, testing testingRepo.Testing, timeInternal timew.Time, dlUser dlUser.DataLogic) auth.UseCase {
	return &useCase{
		log:          log,
		cfg:          cfg,
		testingRepo:  testing,
		timeInternal: timeInternal,
		dlUser:       dlUser,
	}
}
