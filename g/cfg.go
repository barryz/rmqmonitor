package g

import (
	"encoding/json"
	"github.com/toolkits/file"
	"log"
	"sync"
)

type RabbitConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type FalconConfig struct {
	Api string `json:"api"`
}

type GlobalConfig struct {
	Debug     bool          `json:"debug"`
	Details   bool          `json:"details`
	Hostname  string        `json:"hostname"`
	Version   string        `json:"version"`
	Batchsize int           `json:"batchsize"`
	Interval  int64         `json:"interval"`
	Rabbit    *RabbitConfig `json:"rabbitmq"`
	Falcon    *FalconConfig `json:"falcon"`
	Ignores   []string      `json:"ignore_queue"`
	Qrunning  []string      `json:"qrunning"`
}

var (
	config *GlobalConfig
	lock   = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Println("use -c to specify configuration file")
	}

	var c GlobalConfig

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	config = &c

	log.Println("read config file:", cfg, "successfully")
}
