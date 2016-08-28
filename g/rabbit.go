package g

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"fmt"
	"net"
	"time"
)

func GetHost() string {
	host := Config().Rabbit.Host
	if host != "" {
		return host
	}

	host, err := os.Hostname()
	if err != nil {
		log.Println("ERROR: os.Hostname() fail", err)
		return "127.0.0.1"
	}
	return host
}

func GetApiUrl(service string) string {
	host := GetHost()
	port := Config().Rabbit.Port
	api_url := "http://" + host + ":" + strconv.Itoa(port) + "/api/" + service
	return api_url
}

func RabbitApi(service string) ([]byte, error) {
	url := GetApiUrl(service)
	user := Config().Rabbit.User
	password := Config().Rabbit.Password

	// set connect/get/resp timeout
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second * 3)
				if err != nil {
					log.Println("ERROR: dail timeout", err)
					return nil, err
				}
				return c, nil

			},
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 2,
		},
	}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "application/json")
	request.SetBasicAuth(user, password)
	response, err := client.Do(request)
	if err != nil {
		return []byte(""), fmt.Errorf("ERROR: Call rabbit api fail")
	}

	defer response.Body.Close()

	result_code := response.StatusCode
	switch result_code {
	case 200:
		body, _ := ioutil.ReadAll(response.Body)
		return body, nil
	case 401:
		return []byte(""), fmt.Errorf("WARNING: Call rabbitmq rest api auth fail")
	default:
		return []byte(""), fmt.Errorf("ERROR: Unknown error")
	}
}
