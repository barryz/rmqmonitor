package httplib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PostJSONRequest new http request for post json
type PostJSONRequest struct {
	req *http.Request
}

// SetHeader set new header to PostJSONRequest
func (req *PostJSONRequest) SetHeader(k, v string) {
	req.req.Header.Set(k, v)
}

// SetBasicAuth set a basic authentication to PostJSONRequest
func (req *PostJSONRequest) SetBasicAuth(u, p string) {
	req.req.SetBasicAuth(u, p)
}

// DoRequest send a request
func (req *PostJSONRequest) DoRequest() (response []byte, err error) {
	var resp *http.Response
	client := &http.Client{}
	resp, err = client.Do(req.req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		err = fmt.Errorf("get invalid response status code %d", resp.StatusCode)
		return
	}

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// PostJSONWithParams PostJson with some headers or params
func PostJSONWithParams(url string, v interface{}) (newreq *PostJSONRequest, err error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bs))
	if err != nil {
		return
	}

	newreq = &PostJSONRequest{req: req}

	return
}

// PostJSON Post Json object
func PostJSON(url string, v interface{}) (response []byte, err error) {
	var bs []byte
	bs, err = json.Marshal(v)
	if err != nil {
		return
	}

	bf := bytes.NewBuffer(bs)

	var resp *http.Response
	resp, err = http.Post(url, "application/json", bf)
	if err != nil {
		return
	}

	if resp.Body != nil {
		defer resp.Body.Close()
		response, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
	}

	if resp.StatusCode != 200 {
		err = errors.New("status code not equals 200")
	}

	return
}
