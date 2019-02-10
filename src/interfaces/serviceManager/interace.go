package serviceManager

import (
	"interfaces/service"
)

const RUN_SERVICE_CMD = 0x01

func RunService(cmdChan *service.CtrlChanel, serviceSOName string) {
	*cmdChan <- service.CtrlChanelCmd{
		RUN_SERVICE_CMD,
		[]string{serviceSOName},
	}
}
