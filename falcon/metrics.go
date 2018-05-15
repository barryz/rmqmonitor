package falcon

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/barryz/rmqmonitor/funcs"
	"github.com/barryz/rmqmonitor/g"
)

var (
	statsDB = g.NewStatsDB()
)

const (
	overviewPrefix = "rabbitmq.overview."
	queuePrefix    = "rabbitmq.queue."
	exchangePrefix = "rabbitmq.exchange."
)

// MetaData meta data
type MetaData struct {
	Endpoint    string      `json:"endpoint"`
	Metric      string      `json:"metric"`
	Value       interface{} `json:"value"`
	CounterType string      `json:"counterType"`
	Tags        string      `json:"tags"`
	Timestamp   int64       `json:"timestamp"`
	Step        int64       `json:"step"`
}

// NewMetric create an new metric
func NewMetric(name string, value interface{}, tags string) *MetaData {
	host := g.GetHost()
	return &MetaData{
		Metric:      name,
		Endpoint:    host,
		CounterType: fmt.Sprintf("GAUGE"),
		Tags:        tags,
		Timestamp:   time.Now().Unix(),
		Step:        g.Config().Interval,
		Value:       value,
	}
}

func (m *MetaData) String() string {
	s := fmt.Sprintf("MetaData Metric:%s Endpoint:%s Value:%v CounterType:%s Tags:%s Timestamp:%d Step:%d",
		m.Metric, m.Endpoint, m.Value, m.CounterType, m.Tags, m.Timestamp, m.Step)
	return s
}

// SetValue value setter
func (m *MetaData) SetValue(v interface{}) {
	m.Value = v
}

func trimFloat(s float64) float64 {
	if s, err := strconv.ParseFloat(fmt.Sprintf("%.3f", s), 64); err == nil {
		return s
	}
	return s
}

func calcPercentage(l, t int64) (pct float64) {
	if t == 0 {
		return
	}
	pct = float64(l) / float64(t) * 100.00
	pct = trimFloat(pct)
	return
}

func qStats(s string) int64 {
	var aliveQueue = g.Config().Qrunning
	for _, i := range aliveQueue {
		if strings.Contains(strings.ToLower(s), i) {
			return 1
		}
	}
	return 0
}

func isAliveness(s string) int64 {
	switch s {
	case "ok":
		return 1
	default:
		return 0
	}
}

func partitions(s []string) int64 {
	switch len(s) {
	case 0:
		return 1
	default:
		return 0
	}
}

func consumerUtil(c interface{}) float64 {
	if vv, ok := c.(float64); ok {
		return trimFloat(vv * 100.00)
	} else if _, ok := c.(bool); ok {
		return 0.0
	} else if _, ok := c.(string); ok {
		return 0.0
	}
	return 0.0
}

func updateCurrentStatsDB(db string) {
	statsDB.SetCurrentLocate(db)
}

// GetCurrentStatsDB get current stats management database
func GetCurrentStatsDB() *g.StatsDB {
	return statsDB
}

// handleJudge
func handleJudge() (data []*MetaData) {
	data = make([]*MetaData, 0)
	nd, err := funcs.GetNode()
	if err != nil {
		log.Println(err.Error())
		return
	}

	data = append(data, NewMetric(overviewPrefix+"ioReadawait", nd.Rawait, ""))    // io_read_avg_wait_time
	data = append(data, NewMetric(overviewPrefix+"ioWriteawait", nd.Wawait, ""))   // io_write_avg_wait_time
	data = append(data, NewMetric(overviewPrefix+"ioSyncawait", nd.Syncawait, "")) // io_sync_avg_wait_time
	data = append(data, NewMetric(overviewPrefix+"memConnreader", nd.ConnectionReaders, ""))
	data = append(data, NewMetric(overviewPrefix+"memConnwriter", nd.ConnectionWriters, ""))
	data = append(data, NewMetric(overviewPrefix+"memConnchannels", nd.ConnectionChannels, ""))
	data = append(data, NewMetric(overviewPrefix+"memMnesia", nd.Mnesia, ""))
	data = append(data, NewMetric(overviewPrefix+"memMgmtdb", nd.MgmtDB, ""))
	data = append(data, NewMetric(overviewPrefix+"memPlugins", nd.Plugins, ""))
	data = append(data, NewMetric(overviewPrefix+"memMsgidx", nd.MsgIndex, ""))
	data = append(data, NewMetric(overviewPrefix+"memBinary", nd.Binary, ""))
	data = append(data, NewMetric(overviewPrefix+"memAlarm", nd.MemAlarmStatus(), ""))
	data = append(data, NewMetric(overviewPrefix+"diskAlarm", nd.DiskAlarmStatus(), ""))
	data = append(data, NewMetric(overviewPrefix+"fdUsedPct", calcPercentage(nd.FdUsed, nd.FdTotal), ""))
	data = append(data, NewMetric(overviewPrefix+"memUsedPct", calcPercentage(nd.MemUsed, nd.MemLimit), ""))
	data = append(data, NewMetric(overviewPrefix+"socketUsedPct", calcPercentage(nd.SocketsUsed, nd.SocketsTotal), ""))
	data = append(data, NewMetric(overviewPrefix+"erlProcsUsedPct", calcPercentage(nd.ErlProcUsed, nd.ErlProcTotal), "")) //消费生产比
	data = append(data, NewMetric(overviewPrefix+"runQueue", nd.RunQueues, ""))
	data = append(data, NewMetric(overviewPrefix+"isPartition", partitions(nd.Partitions), "")) // 是否发生网络分区

	currentNode := "rabbit@" + g.GetHost()
	ov, err := funcs.GetOverview()
	if err != nil {
		log.Println(err.Error())
		return
	}

	updateCurrentStatsDB(ov.StatisticsDbNode)

	// RabbitMQ Version Compatibility: (<= 3.6.x)
	if ov.StatisticsDbNode == currentNode || len(ov.StatisticsDbNode) == 0 {
		channelCost, err := funcs.GetChannelCost()
		if err != nil {
			log.Println(err.Error())
			return
		}

		aliveness, err := funcs.GetAlive()
		if err != nil {
			log.Printf("get aliveness api failed due to %s", err.Error())
			return
		}

		queues, err := funcs.GetQueues()
		if err != nil {
			log.Printf("get queue api failed due to %s", err.Error())
			return
		}

		exchs, err := funcs.GetExchanges()
		if err != nil {
			log.Printf("get exchange api failed due to %s", err.Error())
			return
		}

		data = append(data, NewMetric(overviewPrefix+"queuesTotal", ov.Queues, "")) // 队列总数
		data = append(data, NewMetric(overviewPrefix+"channelsTotal", ov.Channels, ""))
		data = append(data, NewMetric(overviewPrefix+"connectionsTotal", ov.Connections, ""))
		data = append(data, NewMetric(overviewPrefix+"consumersTotal", ov.Consumers, ""))
		data = append(data, NewMetric(overviewPrefix+"exchangesTotal", ov.Exchanges, ""))
		data = append(data, NewMetric(overviewPrefix+"msgsTotal", ov.MsgsTotal, ""))
		data = append(data, NewMetric(overviewPrefix+"msgsReadyTotal", ov.MsgsReadyTotal, ""))
		data = append(data, NewMetric(overviewPrefix+"msgsUnackTotal", ov.MsgsUnackedTotal, ""))
		data = append(data, NewMetric(overviewPrefix+"deliverTotal", ov.DeliverGet, ""))
		data = append(data, NewMetric(overviewPrefix+"publishTotal", ov.Publish, ""))
		data = append(data, NewMetric(overviewPrefix+"redeliverTotal", ov.Redeliver, ""))
		data = append(data, NewMetric(overviewPrefix+"statsDbEvent", ov.StatsDbEvents, "")) //统计数据库事件数
		data = append(data, NewMetric(overviewPrefix+"deliverRate", ov.DeliverGetRates.Rate, ""))
		data = append(data, NewMetric(overviewPrefix+"publishRate", ov.PublishRates.Rate, ""))
		data = append(data, NewMetric(overviewPrefix+"confirmRate", ov.ConfirmRates.Rate, ""))
		data = append(data, NewMetric(overviewPrefix+"redeliverRate", ov.RedeliverRates.Rate, ""))
		data = append(data, NewMetric(overviewPrefix+"ackRate", ov.AckRates.Rate, ""))
		data = append(data, NewMetric(overviewPrefix+"getChannelCost", channelCost, "")) // 获取channel耗时
		data = append(data, NewMetric(overviewPrefix+"dpRatio", calcPercentage(int64(ov.DeliverGetRates.Rate), int64(ov.PublishRates.Rate)), ""))
		data = append(data, NewMetric(overviewPrefix+"isAlive", isAliveness(aliveness.Status), "")) // 读写判断
		data = append(data, NewMetric(overviewPrefix+"isUp", 1, ""))

		for _, q := range queues {
			tags := fmt.Sprintf("name=%s,vhost=%s", q.Name, q.Vhost)
			data = append(data, NewMetric(queuePrefix+"messages", q.Messages, tags))
			data = append(data, NewMetric(queuePrefix+"messages_ready", q.MessagesReady, tags))
			data = append(data, NewMetric(queuePrefix+"messages_unacked", q.MessagesUnacked, tags))
			data = append(data, NewMetric(queuePrefix+"deliver_get", q.DeliverGet.Rate, tags))
			data = append(data, NewMetric(queuePrefix+"publish", q.Publish.Rate, tags))
			data = append(data, NewMetric(queuePrefix+"redeliver", q.Redeliver.Rate, tags))
			data = append(data, NewMetric(queuePrefix+"ack", q.Ack.Rate, tags))
			data = append(data, NewMetric(queuePrefix+"memory", q.Memory, tags))
			data = append(data, NewMetric(queuePrefix+"consumers", q.Consumers, tags))
			data = append(data, NewMetric(queuePrefix+"consumer_utilisation", consumerUtil(q.ConsumerUtil), tags))
			data = append(data, NewMetric(queuePrefix+"status", qStats(q.Status), tags))
			data = append(data, NewMetric(queuePrefix+"dpratio", calcPercentage(int64(q.DeliverGet.Rate), int64(q.Publish.Rate)), tags))
		}

		for _, e := range exchs {
			tags := fmt.Sprintf("name=%s,vhost=%s", e.Name, e.VHost)
			data = append(data, NewMetric(exchangePrefix+"publish_in", e.MsgStats.PublishInRate.Rate, tags))
			data = append(data, NewMetric(exchangePrefix+"publish_out", e.MsgStats.PublishOutRate.Rate, tags))
			data = append(data, NewMetric(exchangePrefix+"confirm", e.MsgStats.ConfirmRate.Rate, tags))
		}
	}

	return
}

func handleSickRabbit() (data []*MetaData) {
	data = make([]*MetaData, 0)
	data = append(data, NewMetric(overviewPrefix+"isUp", 0, ""))
	return
}

// Collector collect metrics
func Collector() {
	var m []*MetaData

	if !funcs.CheckAlive() {
		log.Println("[ERROR]: Can not connect to rabbit.")
		m = handleSickRabbit()
	} else {
		m = handleJudge()
	}

	log.Printf("[INFO]: send to %s, size: %d.", g.Config().Falcon.API, len(m))
	// log for debug
	if g.Config().Debug {
		for _, m := range m {
			log.Println(m.String())
		}
	}
	sendDatas(m)
}
