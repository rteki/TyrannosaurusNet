package main

import (
	"fmt"
	"interfaces/serviceManager"
)

func (sm *ServiceManager) service() {
	for {
		select {
		case cmd := <-sm.ctrl:
			{
				switch cmd.Id {
				case serviceManager.RUN_SERVICE_CMD:
					{
						sm.cmdRunService(cmd.Params[0])
					}
				}
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

func (sm *ServiceManager) cmdRunService(serviceSOName string) {
	fmt.Println("Run Service", serviceSOName)
}
