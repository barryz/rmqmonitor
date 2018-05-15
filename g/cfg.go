package g

import (
	"log"
	"os"
	"sync"

	"github.com/toolkits/file"
	"gopkg.in/yaml.v2"
)

// EnableConfig configs which can be used
type EnableConfig struct {
	Collect   bool `yaml:"collect"`
	LogRotate bool `yaml:"log_rotate"`
	Witch     bool `yaml:"witch"`
}

// RabbitConfig ...
type RabbitConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// FalconConfig ...
type FalconConfig struct {
	API string `yaml:"api"`
}

// HTTPConfig ...
type HTTPConfig struct {
	ConnTimeout int `yaml:"conn_timeout"`
	RespTimeout int `yaml:"response_timeout"`
}

// SchedulerConfig ...
type SchedulerConfig struct {
	LogRotate string `yaml:"log_rotate"`
}

// WitchConfig Program Config ...
type WitchConfig struct {
	ListenAddr string            `yaml:"listen"`
	Control    string            `yaml:"control"`
	Service    string            `yaml:"service"`
	Process    string            `yaml:"process"`
	Command    string            `yaml:"command"`
	PidFile    string            `yaml:"pid_file"`
	Auth       map[string]string `yaml:"auth"`
}

// GlobalConfig ...
type GlobalConfig struct {
	Debug     bool             `yaml:"debug"`
	Details   bool             `yaml:"details"`
	Hostname  string           `yaml:"hostname"`
	Batchsize int              `yaml:"batchsize"`
	Interval  int64            `yaml:"interval"`
	Rabbit    *RabbitConfig    `yaml:"rabbitmq"`
	Falcon    *FalconConfig    `yaml:"falcon"`
	HTTP      *HTTPConfig      `yaml:"http"`
	Cron      *SchedulerConfig `yaml:"scheduler"`
	Enabled   *EnableConfig    `yaml:"enabled"`
	Ignores   []string         `yaml:"ignore_queue"`
	Qrunning  []string         `yaml:"qrunning"`
	Witch     *WitchConfig     `yaml:"witch"`
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

	err = yaml.Unmarshal([]byte(configContent), &c)
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
