package funcs

import (
	"encoding/json"
	"github.com/barryz/rmqmonitor/g"
	"log"
)

type Rate struct {
	Rate float64 `json:"rate"`
}

type MsgStat struct {
	Publish              int64 `json:"publish"`
	Ack                  int64 `json:"ack"`
	Deliver_get          int64 `json:"deliver_get"`
	Redeliver            int64 `json:"redeliver"`
	Confirm              int64 `json:"confirm"`
	Deliver              int64 `json:"deliver"`
	Deliver_no_ack       int64 `json:"deliver_no_ack"`
	Get                  int64 `json:"get"`
	Get_no_ack           int64 `json:"get_no_ack"`
	Publish_Rates        Rate  `json:"publish_details"`
	Deliver_get_Rates    Rate  `json:"deliver_get_details"`
	Ack_Rates            Rate  `json:"ack_details"`
	Confirm_Rates        Rate  `json:"confirm_details"`
	Redeliver_Rates      Rate  `json:"redeliver_details"`
	Deliver_Rates        Rate  `json:"deliver_details"`
	Deliver_no_ack_Rates Rate  `json:"deliver_no_ack_details"`
	Get_no_ack_Rates     Rate  `json:"get_no_ack_details"`
	Get_Rates            Rate  `json:"get_details"`
}

type QueueTotal struct {
	MsgsTotal        int64 `json:"messages"`
	MsgsReadyTotal   int64 `json:"messages_ready"`
	MsgsUnackedTotal int64 `json:"messages_unacknowledged"`
}

type ObjectTotal struct {
	Consumers   int64 `json:"consumers"`
	Queues      int64 `json:"queues"`
	Exchanges   int64 `json:"exchanges"`
	Connections int64 `json:"connections"`
	Channels    int64 `json:"channels"`
}

type OverView struct {
	MsgStat       `json:"message_stats"`
	QueueTotal    `json:"queue_totals"`
	ObjectTotal   `json:"object_totals"`
	StatsDbEvents int `json:"statistics_db_event_queue"`
}

func GetOverview() *OverView {
	var service string = "overview"
	var result OverView
	res, err := g.RabbitApi(service)
	if err != nil {
		log.Println(err)
		return &result
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		log.Println("ERROR: unmarshal json data fail, ", err)
		return &result
	}

	return &result
}
