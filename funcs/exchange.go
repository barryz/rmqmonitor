package funcs

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/barryz/rmqmonitor/g"
)

// VHostName name of vHost
type VHostName struct {
	Name string `json:"name"`
}

// ExchangeInfo information for exchange
type ExchangeInfo struct {
	Name     string `json:"name"`
	VHost    string `json:"vhost"`
	MsgStats struct {
		Confirm    int64 `json:"confirm"`
		PublishIn  int64 `json:"publish_in"`
		PublishOut int64 `json:"publish_out"`

		ConfirmRate struct {
			Rate float64 `json:"rate"`
		} `json:"confirm_details"`

		PublishInRate struct {
			Rate float64 `json:"rate"`
		} `json:"publish_in_details"`

		PublishOutRate struct {
			Rate float64 `json:"rate"`
		} `json:"publish_out_details"`
	} `json:"message_stats"`
}

func getVHosts() (vl []string, err error) {
	service := "vhosts"
	res, err := g.RabbitAPI(service)
	if err != nil {
		err = fmt.Errorf("[ERROR]: get rabbitmq vhost info fail due to %s", err.Error())
		return
	}

	var vs []*VHostName
	err = json.Unmarshal(res, &vs)
	if err != nil {
		err = fmt.Errorf("[ERROR]: unmarshal rabbitmq vhost json data fail due to %s", err.Error())
		return
	}

	vl = make([]string, len(vs))
	for _, v := range vs {
		vl = append(vl, urlEncode(v.Name))
	}
	return
}

// GetExchanges get all exchanges
func GetExchanges() (exchs []*ExchangeInfo, err error) {
	vhosts, err := getVHosts()
	if err != nil {
		return
	}

	exchs = make([]*ExchangeInfo, 0)
	for _, v := range vhosts {
		var (
			es   []*ExchangeInfo
			err1 error
		)
		service := fmt.Sprintf("exchanges/%s", v)
		res, err1 := g.RabbitAPI(service)
		if err1 != nil {
			err = err1
			return
		}

		err1 = json.Unmarshal(res, &es)
		if err1 != nil {
			err = err1
			return
		}

		for _, e := range es {
			if e.Name == "" {
				e.Name = "DEFAULT_EXCHANGE"
			}

			if strings.Contains(e.Name, "amq.") {
				continue
			}
			exchs = append(exchs, e)
		}
	}

	return
}
