package main

import (
	"fmt"
	"interfaces/service"
)

type ServiceManager struct {
	serviceName     string
	managedServices []*service.Service
	stopchan        service.KillChanel
	ctrl            service.CtrlChanel
	ctrls           map[string]func(service.CtrlChanelParams)
}

func (sm *ServiceManager) Run() {
	sm.stopchan = make(service.KillChanel)
	sm.ctrl = make(service.CtrlChanel)
	go sm.service()
}

func (sm *ServiceManager) Stop() {
	close(sm.stopchan)
	fmt.Println("Services Manager stoped.")
}

func (sm *ServiceManager) GetName() string {
	return sm.serviceName
}

func (sm *ServiceManager) GetCtrl() *service.CtrlChanel {
	return &sm.ctrl
}

func main() {}

func constructor() service.Service {
	sm := new(ServiceManager)
	sm.serviceName = "ServiceManager"
	sm.ctrls = make(map[string]func(service.CtrlChanelParams))
	sm.ctrls["runService"] = sm.runService
	return sm
}

var New service.ServiceConstructor = constructor
