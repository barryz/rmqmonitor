package cron


import (
	"testing"

	"github.com/barryz/rmqmonitor/g"
)

func Test_LogRotateStart(t *testing.T) {
	g.ParseConfig("../cfg.json")

	go CronStart()

	select {}
}

