package falcon

import (
	"testing"

	"github.com/barryz/rmqmonitor/g"
)

func TestGetStatsDB(t *testing.T) {
	g.ParseConfig("../cfg.json.example")

	//ov, err := funcs.GetOverview()
	//if err != nil {
	//	t.Fatalf("%s", err.Error())
	//}
	//
	//updateCurrentStatsDB(ov.StatisticsDbNode)
	//
	//stats := GetCurrentStatsDB()
	//fmt.Printf("%s --- %s \n", stats.CurrentLocate, stats.PreviousLocate)
}
