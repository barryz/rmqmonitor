package falcon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/barryz/rmqmonitor/g"
)

func sendData(data []*MetaData) (resp []byte, err error) {
	debug := g.Config().Debug
	js, err := json.Marshal(data)
	if err != nil {
		return
	}

	if debug {
		log.Printf("agent api received %d metrics", len(data))
	}

	res, err := http.Post(g.Config().Falcon.API, "Content-Type: application/json", bytes.NewBuffer(js))
	if err != nil {
		err = fmt.Errorf("[ERROR]: sent data to falcon agent api fail due to %s", err.Error())
		return
	}

	defer res.Body.Close()

	return
}

func sendDatas(m []*MetaData) {
	// batch-size specified.
	limit, lens := g.Config().Batchsize, len(m)
	if lens >= limit {
		offset := lens % limit
		for i := 0; i <= lens-1; i += limit {
			if (i + limit - 1) >= lens {
				_, err := sendData(m[i:(offset + i - 1)])
				if err != nil {
					log.Println(err.Error())
					break
				}
			} else {
				_, err := sendData(m[i:(limit + i - 1)])
				if err != nil {
					log.Println(err.Error())
					break
				}
			}
		}
	} else {
		_, err := sendData(m)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
