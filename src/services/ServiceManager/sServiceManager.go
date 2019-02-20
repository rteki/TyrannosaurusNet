package main

import (
	"fmt"
	"interfaces/service"
	"time"
)

func (sm *ServiceManager) service() {
	for {
		select {
		case cmd := <-sm.ctrl:
			{
				sm.ctrls[cmd.Id](cmd.Params)
			}
		default:
			{
			}
		case <-sm.stopchan:
			{
				fmt.Print("Closing....")
				return
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func (sm *ServiceManager) runService(args service.CtrlChanelParams) {
	s := service.LoadServiceFromPlugin(string(args[0]))
	sm.managedServices = append(sm.managedServices, s)

	s.Run()
}
