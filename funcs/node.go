package funcs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/barryz/rmqmonitor/g"
)

// MemStats ...
type MemStats struct {
	Total              int64 `json:"total"`
	ConnectionReaders  int64 `json:"connection_readers"`
	ConnectionWriters  int64 `json:"connection_writers"`
	ConnectionChannels int64 `json:"connection_channels"`
	ConnectionOther    int64 `json:"connection_other"`
	QueueProcs         int64 `json:"queue_procs"`
	QueueSlaveProcs    int64 `json:"queue_slave_procs"`
	Plugins            int64 `json:"plugins"`
	Mnesia             int64 `json:"mnesia"`
	MgmtDB             int64 `json:"mgmt_db"`
	MsgIndex           int64 `json:"msg_index"`
	Code               int64 `json:"code"`
	Atom               int64 `json:"atom"`
	Binary             int64 `json:"binary"`
}

// NodeStats ...
type NodeStats struct {
	MemStats     `json:"memory"`
	Partitions   []string `json:"partitions"`
	Rawait       float64  `json:"io_read_avg_time"`
	Wawait       float64  `json:"io_write_avg_time"`
	Syncawait    float64  `json:"io_sync_avg_time"`
	MemUsed      int64    `json:"mem_used"`
	MemLimit     int64    `json:"mem_limit"`
	SocketsUsed  int64    `json:"sockets_used"`
	SocketsTotal int64    `json:"sockets_total"`
	FdUsed       int64    `json:"fd_used"`
	FdTotal      int64    `json:"fd_total"`
	ErlProcUsed  int64    `json:"proc_used"`
	ErlProcTotal int64    `json:"proc_total"`
	RunQueues    int64    `json:"run_queue"`
	MemAlarm     bool     `json:"mem_alarm"`
	DiskAlarm    bool     `json:"disk_free_alarm"`
}

// MemAlarmStatus 内存告警指标
func (n *NodeStats) MemAlarmStatus() int {
	if n.MemAlarm {
		return 0
	}
	return 1
}

// DiskAlarmStatus 磁盘告警指标
func (n *NodeStats) DiskAlarmStatus() int {
	if n.DiskAlarm {
		return 0
	}
	return 1
}

// GetNode ...
func GetNode() (n *NodeStats, err error) {
	host := g.GetHost()
	if g.Config().Debug {
		log.Printf("[INFO]: Get hostname %s.", host)
	}

	service := "nodes/rabbit@" + host + "?memory=true"
	// service := "nodes/rabbit@" + "vm-test-barryz" + "?memory=true"
	res, err := g.RabbitAPI(service)
	if err != nil {
		err = fmt.Errorf("[ERROR]: get rabbitmq node info fail due to %s", err.Error())
		return
	}

	err = json.Unmarshal(res, &n)
	if err != nil {
		err = fmt.Errorf("[ERROR]: unmarshal rabbitmq node info json data fail due to %s", err.Error())
		return
	}
	return
}
