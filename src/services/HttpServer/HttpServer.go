package main

import (
	"fmt"
	"interfaces/service"
)

type HttpServer struct {
	serviceName string
	stopchan    service.KillChanel
	ctrl        service.CtrlChanel
	ctrls       map[string]func(service.CtrlChanelParams)
}

func (hp *HttpServer) Run() {
	hp.stopchan = make(service.KillChanel)
	hp.ctrl = make(service.CtrlChanel)
	go hp.service()
}

func (hp *HttpServer) Stop() {
	close(hp.stopchan)
	fmt.Println("Services Manager stoped.")
}

func (hp *HttpServer) GetName() string {
	return hp.serviceName
}

func (hp *HttpServer) GetCtrl() *service.CtrlChanel {
	return &hp.ctrl
}

func main() {}

func constructor() service.Service {
	hp := new(HttpServer)
	hp.serviceName = "HttpServer"
	hp.ctrls = make(map[string]func(service.CtrlChanelParams))
	return hp
}

var New service.ServiceConstructor = constructor
