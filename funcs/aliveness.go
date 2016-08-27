package funcs

import (
	"encoding/json"
	"fmt"
	"rmqmon/g"
)

type AliveNess struct {
	Status string `json:"status"`
}

func GetAlive() (*AliveNess, error) {
	service := "aliveness-test/%2f"
	var result AliveNess

	res, err := g.RabbitApi(service)
	if err != nil {
		return &result, err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return &result, fmt.Errorf("ERROR: unmarshal json data fail")
	}

	return &result, nil
}

func CheckAlive() bool {
	service := "whoami"
	if _, err := g.RabbitApi(service); err == nil {
		return true
	} else {
		return false
	}
}
