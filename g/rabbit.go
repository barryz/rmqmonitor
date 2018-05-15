package g

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

// GetHost get hostname (syscall)
func GetHost() string {
	return Config().Hostname
}

// GetAPIUrl return the RabbitMQ api url
func GetAPIUrl(service string) string {
	port := Config().Rabbit.Port
	apiURL := fmt.Sprintf("http://%s:%s/api/%s", Config().Rabbit.Host, strconv.Itoa(port), service)
	return apiURL
}

// RabbitAPI ...
func RabbitAPI(service string) ([]byte, error) {
	url := GetAPIUrl(service)
	user := Config().Rabbit.User
	connTimeout := Config().HTTP.ConnTimeout
	respTimeout := Config().HTTP.RespTimeout
	password := Config().Rabbit.Password

	// set connect/get/resp timeout
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(connTimeout))
				if err != nil {
					return nil, err
				}
				return c, nil

			},
			DisableKeepAlives:     true, // disable http keepalive
			ResponseHeaderTimeout: time.Second * time.Duration(respTimeout),
		},
	}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(user, password)
	response, err := client.Do(request)
	if err != nil {
		return []byte(""), fmt.Errorf("call rabbit api fail")
	}

	defer response.Body.Close()

	resultCode := response.StatusCode
	switch resultCode {
	case http.StatusOK:
		body, _ := ioutil.ReadAll(response.Body)
		return body, nil
	case http.StatusUnauthorized:
		return []byte(""), fmt.Errorf("call rabbitmq rest api auth fail")
	default:
		return []byte(""), fmt.Errorf("unknown error %d", resultCode)
	}
}
