package main

import (
	testingRepo "github.com/AndreeJait/template-service-go/internal/repository/db/testing"
	testingRepoImpl "github.com/AndreeJait/template-service-go/internal/repository/db/testing/impl"
)

type listRepository struct {
	testingRepo testingRepo.Testing
}

func (a *Api) initRepository() listRepository {

	var list listRepository

	list.testingRepo = testingRepoImpl.NewRepository(a.sqlW, a.log, a.cfg, a.appMode)

	return list
}
