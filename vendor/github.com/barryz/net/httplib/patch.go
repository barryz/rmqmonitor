package httplib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PatchJSONRequest new http request for patch json
type PatchJSONRequest struct {
	req *http.Request
}

// SetHeader set new header to PatchJSONRequest
func (req *PatchJSONRequest) SetHeader(k, v string) {
	req.req.Header.Set(k, v)
}

// SetBasicAuth set a basic authentication to PatchJSONRequest
func (req *PatchJSONRequest) SetBasicAuth(u, p string) {
	req.req.SetBasicAuth(u, p)
}

// DoRequest send a request
func (req *PatchJSONRequest) DoRequest() (response []byte, err error) {
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

// PatchJSONWithParams PostJson with some headers or params
func PatchJSONWithParams(url string, v interface{}) (newreq *PatchJSONRequest, err error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bs))
	if err != nil {
		return
	}

	newreq = &PatchJSONRequest{req: req}

	return
}
