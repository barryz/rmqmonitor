package funcs

import (
	"encoding/json"
	"rmqmon/g"
	"strings"
	"log"
)

type QueueRate struct {
	Rate float64 `json:"rate"`
}

type QueueMsgStat struct {
	Publish     QueueRate `json:"publish_details"`
	Deliver_get QueueRate `json:"deliver_get_details"`
	Ack         QueueRate `json:"ack_details"`
	Redeliver   QueueRate `json:"redeliver_details"`
}

type QueueMap struct {
	Memory          int64   `json:"memory"`
	Messages        int64   `json:"messages"`
	MessagesReady   int64   `json:"messages_ready"`
	MessagesUnacked int64   `json:"messages_unacknowledged"`
	ConsumerUtil    float64 `json:"consumer_utilisation"`
	Consumers       int64   `json:"consumers"`
	Status          string  `json:"state"`
	Name            string  `json:"name"`
	Vhost           string  `json:"vhost"`
	Auto_Delete     bool    `json:"auto_delete"`
	QueueMsgStat    `json:"message_stats"`
}

func filterQueue(q *QueueMap) bool {
	var (
		isignore bool = false
		isad     bool = false
		isvhost  bool = false
	)
	ignores := g.Config().Ignores

	for _, i := range ignores {
		if strings.Contains(strings.ToLower(q.Name), i) {
			isignore = true
		} else {
			continue
		}
	}

	if q.Auto_Delete {
		isad = true
	}

	if q.Vhost == "/" {
		isvhost = true
	}

	if isignore || isad || isvhost {
		return true
	} else {
		return false
	}
}

func GetQueues() *[]QueueMap {
	var (
		queues    []QueueMap
		newqueues []QueueMap
	)

	service := "queues"
	res, err := g.RabbitApi(service)
	if err != nil {
		log.Println(err)
		return &newqueues
	}
	if err := json.Unmarshal(res, &queues); err != nil {
		log.Println("ERROR: unmarshal queue file fail")
		return &newqueues
	}

	for _, q := range queues {
		if !filterQueue(&q) {
			newqueues = append(newqueues, q)
		}
	}

	return &newqueues
}
