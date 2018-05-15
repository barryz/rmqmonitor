package funcs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/barryz/rmqmonitor/g"

	"github.com/streadway/amqp"
)

// Rate ...
type Rate struct {
	Rate float64 `json:"rate"`
}

// MsgStat ...
type MsgStat struct {
	Publish           int64 `json:"publish"`
	Ack               int64 `json:"ack"`
	DeliverGet        int64 `json:"deliver_get"`
	Redeliver         int64 `json:"redeliver"`
	Confirm           int64 `json:"confirm"`
	Deliver           int64 `json:"deliver"`
	DeliverNoAck      int64 `json:"deliver_no_ack"`
	Get               int64 `json:"get"`
	GetNoAck          int64 `json:"get_no_ack"`
	PublishRates      Rate  `json:"publish_details"`
	DeliverGetRates   Rate  `json:"deliver_get_details"`
	AckRates          Rate  `json:"ack_details"`
	ConfirmRates      Rate  `json:"confirm_details"`
	RedeliverRates    Rate  `json:"redeliver_details"`
	DeliverRates      Rate  `json:"deliver_details"`
	DeliverNoAckRates Rate  `json:"deliver_no_ack_details"`
	GetNoAckRates     Rate  `json:"get_no_ack_details"`
	GetRates          Rate  `json:"get_details"`
}

// QueueTotal ...
type QueueTotal struct {
	MsgsTotal        int64 `json:"messages"`
	MsgsReadyTotal   int64 `json:"messages_ready"`
	MsgsUnackedTotal int64 `json:"messages_unacknowledged"`
}

// ObjectTotal ...
type ObjectTotal struct {
	Consumers   int64 `json:"consumers"`
	Queues      int64 `json:"queues"`
	Exchanges   int64 `json:"exchanges"`
	Connections int64 `json:"connections"`
	Channels    int64 `json:"channels"`
}

// OverView ...
type OverView struct {
	MsgStat          `json:"message_stats"`
	QueueTotal       `json:"queue_totals"`
	ObjectTotal      `json:"object_totals"`
	StatsDbEvents    int    `json:"statistics_db_event_queue"`
	StatisticsDbNode string `json:"statistics_db_node"`
}

// GetOverview ...
func GetOverview() (result *OverView, err error) {
	service := "overview"
	res, err := g.RabbitAPI(service)
	if err != nil {
		err = fmt.Errorf("[ERROR]: get rabbitmq overview info fail due to %s", err.Error())
		return
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		err = fmt.Errorf("[ERROR]: unmarshal rabbitmq overview json data fail due to %s", err.Error())
		return
	}

	return
}

// GetChannelCost time cost of getting channel
func GetChannelCost() (getChannelCost float64, err error) {
	uri := fmt.Sprintf("amqp://%s:%s@127.0.0.1:5672//", g.Config().Rabbit.User, g.Config().Rabbit.Password)

	conn, err := amqp.Dial(uri)
	if err != nil {
		err = fmt.Errorf("[ERROR]: get amqp connection fail due to %s", err.Error())
		return
	}

	timeToStart := time.Now()
	ch, err := conn.Channel()
	getChannelCost = time.Now().Sub(timeToStart).Seconds() * 1000
	if err != nil {
		err = fmt.Errorf("[ERROR]: get amqp channel fail due to %s", err.Error())
		return
	}

	ch.Close()
	conn.Close()

	return
}
