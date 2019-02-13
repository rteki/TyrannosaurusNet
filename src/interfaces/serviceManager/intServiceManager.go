package serviceManager

import (
	"interfaces/service"
)

func RunService(cmdChan *service.CtrlChanel, serviceSOName string) {
	*cmdChan <- service.CtrlChanelCmd{
		"runService",
		service.CtrlChanelParams{service.CtrlChanelParam(serviceSOName)},
	}
}
