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

func sendDatas(m []*MetaData) {
	//  根据batchsize发送metrics
	limit, lens := g.Config().Batchsize, len(m)
	if lens >= limit {
		offset := lens % limit
		for i := 0; i <= lens - 1; i += limit {
			if (i + limit - 1) >= lens {
				_, err := sendData(m[i:(offset + i - 1)])
				if err != nil {
					log.Println("ERROR:", err)
					break
				}
			} else {
				_, err := sendData(m[i:(limit + i - 1)])
				if err != nil {
					log.Println("ERROR:", err)
					break
				}
			}
		}
	} else {
		_, err := sendData(m)
		if err != nil {
			log.Println("ERROR:", err)
		}
	}
}
