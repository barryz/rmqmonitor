package falcon

import (
	"fmt"
	"testing"

	"github.com/barryz/rmqmonitor/funcs"
	"github.com/barryz/rmqmonitor/g"
)

func Test_GetStatsDB(t *testing.T) {
	g.ParseConfig("../cfg.json")

	ov, err := funcs.GetOverview()
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	updateCurrentStatsDB(ov.StatisticsDbNode)

	stats := GetCurrentStatsDB()
	fmt.Printf("%s: %s --- %s \n", stats.CurrentLocate, stats.LastUpdateTime, stats.PreviousLocate)
}
