package falcon

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rmqmon/g"
	"log"
)

func sendData(data []*MetaData) ([]byte, error) {
	debug := g.Config().Debug
	js, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if debug {
		log.Printf("agent api recieved %d metrics", len(data))
	}

	res, err := http.Post(g.Config().Falcon.Api, "Content-Type: application/json", bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
