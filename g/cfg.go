package g

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/toolkits/file"
)

// EnableConfig configs which can be used
type EnableConfig struct {
	Collect   bool `json:"collect"`
	LogRotate bool `json:"log_rotate"`
	Witch     bool `json:"witch"`
}

// RabbitConfig ...
type RabbitConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// FalconConfig ...
type FalconConfig struct {
	API string `json:"api"`
}

// HTTPConfig ...
type HTTPConfig struct {
	ConnTimeout int `json:"conn_timeout"`
	RespTimeout int `json:"response_timeout"`
}

// SchedulerConfig ...
type SchedulerConfig struct {
	LogRotate string `json:"log_rotate"`
}

// WitchConfig Program Config ...
type WitchConfig struct {
	ListenAddr string            `json:"listen"`
	Control    string            `json:"control"`
	Service    string            `json:"service"`
	Process    string            `json:"process"`
	Command    string            `json:"command"`
	PidFile    string            `json:"pid_file"`
	Auth       map[string]string `json:"auth"`
}

// GlobalConfig ...
type GlobalConfig struct {
	Debug     bool             `json:"debug"`
	Details   bool             `json:"details"`
	Hostname  string           `json:"hostname"`
	Batchsize int              `json:"batchsize"`
	Interval  int64            `json:"interval"`
	Rabbit    *RabbitConfig    `json:"rabbitmq"`
	Falcon    *FalconConfig    `json:"falcon"`
	HTTP      *HTTPConfig      `json:"http"`
	Cron      *SchedulerConfig `json:"scheduler"`
	Enabled   *EnableConfig    `json:"enabled"`
	Ignores   []string         `json:"ignore_queue"`
	Qrunning  []string         `json:"qrunning"`
	Witch     *WitchConfig     `json:"witch"`
}

var (
	config *GlobalConfig
	lock   = new(sync.RWMutex)
)

// Config ...
func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

// ParseConfig ...
func ParseConfig(cfg string) {
	if cfg == "" {
		log.Println("use -c to specify configuration file")
	}

	var c GlobalConfig

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("[ERROR]: read config file:", cfg, "fail:", err)
	}

	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("[ERROR]: read config file:", cfg, "fail:", err)
	}

	if c.Hostname == "" {
		c.Hostname, err = os.Hostname()
		if err != nil {
			log.Fatalln("[ERROR]: get local hostname fail")
			os.Exit(1)
		}
	}

	config = &c

	log.Println("[INFO]: read config file:", cfg, "successfully")
}
