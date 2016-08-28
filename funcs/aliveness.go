package funcs

import (
	"encoding/json"
	"rmqmon/g"
	"log"
)

type AliveNess struct {
	Status string `json:"status"`
}

func GetAlive() *AliveNess {
	service := "aliveness-test/%2f"
	var result AliveNess

	res, err := g.RabbitApi(service)
	if err != nil {
		log.Println(err)
		return &result
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		log.Println("ERROR: unmarshal json data fail, ", err)
		return &result
	}

	return &result
}

func CheckAlive() bool {
	service := "whoami"
	if _, err := g.RabbitApi(service); err == nil {
		return true
	} else {
		return false
	}
}
