package cron

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/barryz/rmqmonitor/g"
	"github.com/barryz/rmqmonitor/utils"

	"github.com/barryz/cron"
)

func logRotateRun() {
	log.Println("[INFO] start to rotate rabbitmq log file ......")
	suffix := fmt.Sprintf(".%s", utils.GetYesterdayDate())
	logRotateCommand := exec.Command("rabbitmqctl", "rotate_logs", suffix)
	output, err := logRotateCommand.CombinedOutput()
	if err != nil {
		log.Printf("[ERROR]: rotate rabbitmq log failed due to %s", err.Error())
		return
	}
	log.Printf("[INFO] rotate rabbitmq log success, %s", string(output))
	return
}

// Start start the cron tab
func Start() {
	logrotateCron := cron.New()
	logrotateCron.AddFuncCC(g.Config().Cron.LogRotate, func() { logRotateRun() }, 1)
	logrotateCron.Start()
}
