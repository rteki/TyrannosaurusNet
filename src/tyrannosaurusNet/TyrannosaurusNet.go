package main

import (
	"interfaces/service"
	"interfaces/serviceManager"
	"sync"
)

func main() {
	sm := service.LoadServiceFromPlugin("./plugins/ServiceManager.so")

	serverStopWG := new(sync.WaitGroup)

	serverStopWG.Add(1)

	sm.Run()

	serviceManager.RunService(sm.GetCtrl(), "plugins/HttpServer.so")

	serverStopWG.Wait()
}
