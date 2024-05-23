package main

import (
	"github.com/AndreeJait/template-service-go/internal/usecase/auth"
	ucAuthImpl "github.com/AndreeJait/template-service-go/internal/usecase/auth/impl"
)

type listUseCase struct {
	ucAuth auth.UseCase
}

func (a *Api) initUseCase() listUseCase {
	var list listUseCase

	listRepo := a.initRepository()
	listDl := a.initDataLogic()

	// auth useCase
	ucAuth := ucAuthImpl.NewUseCase(a.log, a.cfg, listRepo.testingRepo, a.timeLoc, listDl.dlUser)
	list.ucAuth = ucAuth

	return list
}
