package funcs

import (
	"encoding/json"
	"fmt"

	"github.com/barryz/rmqmonitor/g"
)

// Aliveness ...
type Aliveness struct {
	Status string `json:"status"`
}

// GetAlive ...
func GetAlive() (aliveness *Aliveness, err error) {
	service := "aliveness-test/%2f"

	res, err := g.RabbitAPI(service)
	if err != nil {
		err = fmt.Errorf("[ERROR]: get rabbitmq aliveness fail due to %s", err.Error())
		return
	}

	err = json.Unmarshal(res, &aliveness)
	if err != nil {
		err = fmt.Errorf("[ERROR]: unmarshal rabbitmq aliveness json data fail due to %s", err.Error())
		return
	}

	return
}

// CheckAlive ...
func CheckAlive() (ok bool) {
	service := "whoami"
	if _, err := g.RabbitAPI(service); err == nil {
		ok = true
		return
	}

	return
}
