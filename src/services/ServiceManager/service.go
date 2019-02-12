package main

import (
	"fmt"
	"interfaces/service"
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
	}
}

func (sm *ServiceManager) runService(args service.CtrlChanelParams) {
	fmt.Println("Run Service", args[0])
}
