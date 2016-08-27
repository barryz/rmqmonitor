package funcs

import (
	"encoding/json"
	"rmqmon/g"
	"log"
)

type MemStats struct {
	Total               int64 `json:"total"`
	Connection_readers  int64 `json:"connection_readers"`
	Connection_writers  int64 `json:"connection_writers"`
	Connection_channels int64 `json:"connection_channels"`
	Connection_other    int64 `json:"connection_other"`
	Queue_procs         int64 `json:"queue_procs"`
	Queue_slave_procs   int64 `json:"queue_slave_procs"`
	Plugins             int64 `json:"plugins"`
	Mnesia              int64 `json:"mnesia"`
	Mgmt_db             int64 `json:"mgmt_db"`
	Msg_index           int64 `json:"msg_index"`
	Code                int64 `json:"code"`
	Atom                int64 `json:"atom"`
}

type NodeStats struct {
	MemStats     `json:"memory"`
	Partitions   []string `json:"partitions"`
	Rawait       int64    `json:"io_read_avg_time"`
	Wawait       int64    `json:"io_write_avg_time"`
	Syncawait    int64    `json:"io_sync_avg_time"`
	MemUsed      int64    `json:"mem_used"`
	MemLimit     int64    `json:"mem_limit"`
	SocketsUsed  int64    `json:"sockets_used"`
	SocketsTotal int64    `json:"sockets_total"`
	FdUsed       int64    `json:"fd_used"`
	FdTotal      int64    `json:"fd_total"`
	ErlProcUsed  int64    `json:"proc_used"`
	ErlProcTotal int64    `json:"proc_total"`
	RunQueues    int64    `json:"run_queue"`
}

func GetNode() *NodeStats {
	host := g.GetHost()
	if g.Config().Debug {
		log.Printf("INFO: Get hostname %s.", host)
	}

	service := "nodes/rabbit@" + host + "?memory=true"
	var result NodeStats

	res, err := g.RabbitApi(service)
	if err != nil {
		log.Println(err)
		return &result
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		log.Println("ERROR: unmarshal json data fail")
		return &result
	}

	return &result
}
