package main

import (
	dlUser "github.com/AndreeJait/template-service-go/internal/datalogic/user"
	implDlUser "github.com/AndreeJait/template-service-go/internal/datalogic/user/impl"
)

type listDataLogic struct {
	dlUser dlUser.DataLogic
}

func (a *Api) initDataLogic() listDataLogic {
	var list = listDataLogic{}

	listRepo := a.initRepository()
	list.dlUser = implDlUser.NewDataLogic(a.log, a.cfg, listRepo.testingRepo, a.timeLoc)

	return list
}
