package service

import (
	"fmt"
	"plugin"
)

type Service interface {
	Run()
	Stop()
	GetName() string
	GetCtrl() *CtrlChanel
}

type CtrlChanel chan CtrlChanelCmd

type CtrlChanelParam string

type CtrlChanelParams []CtrlChanelParam

type CtrlChanelCmd struct {
	Id     string
	Params CtrlChanelParams
}

type KillChanel chan struct{}

type ServiceConstructor func() Service

func LoadServiceFromPlugin(path string) Service {

	plug, plugErr := plugin.Open(path)

	if plugErr != nil {
		fmt.Printf("Plugin load error:\n%s", plugErr)
	}

	sym, symErr := plug.Lookup("New")

	if symErr != nil {
		fmt.Printf("Symbol load error:\n%s", symErr)
	}

	New := sym.((*ServiceConstructor))

	service := (*New)()

	return service
}
