package funcs

import (
	"testing"

	"github.com/barryz/rmqmonitor/g"
)

func TestGetExchanges(t *testing.T) {
	g.ParseConfig("../cfg.json.example")
	//exchs, err := GetExchanges()
	//if err != nil {
	//	t.Error(err.Error())
	//}
	//for _, e := range exchs {
	//	fmt.Printf("%s rate is %d, vhost is %s\n", e.Name, e.MsgStats.PublishIn, e.VHost)
	//}
}
