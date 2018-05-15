package witch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/barryz/rmqmonitor/falcon"
	"github.com/barryz/rmqmonitor/g"
	"github.com/barryz/rmqmonitor/witch/system"

	"github.com/martini-contrib/render"
)

var (
	// ErrServerError is internal server error.
	ErrServerError = errors.New("Internal Server Error")
	// ErrBadRequest is bad request error.
	ErrBadRequest = errors.New("Bad Request")
)

func sysAction(control *system.SysController, req *http.Request, r render.Render) {
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("[ERROR] Read request body error: %s", err)
		r.JSON(http.StatusInternalServerError, ErrServerError)
		return
	}
	log.Printf("[INFO] Request action: %s", bs)
	action := &system.Action{}
	if err := json.Unmarshal(bs, action); err != nil {
		log.Printf("[WARN] Invalid action format: %s", err)
		r.JSON(http.StatusBadRequest, ErrBadRequest)
		return
	}
	r.JSON(http.StatusOK, control.Handle(action))
}

func statsAction(control *system.StatsController, req *http.Request, r render.Render) {
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("[ERROR] Read request body error: %s", err)
		r.JSON(http.StatusInternalServerError, ErrServerError)
		return
	}
	log.Printf("[INFO] Request stats action: %s", bs)
	action := &system.Action{}
	if err := json.Unmarshal(bs, action); err != nil {
		log.Printf("[WARN] Invalid action format: %s", err)
		r.JSON(http.StatusBadRequest, ErrBadRequest)
		return
	}
	r.JSON(http.StatusOK, control.Handle(action))

}

func procForceStop(req *http.Request, r render.Render) {
	if req.Method != "GET" {
		r.JSON(http.StatusMethodNotAllowed, ErrBadRequest)
		return
	}
	proc := g.Config().Witch.Process
	args := fmt.Sprintf("pgrep %s|xargs skill -9", proc)
	_, err := system.ExecCommand("bash", []string{"-c", args})
	if err != nil {
		r.JSON(http.StatusServiceUnavailable, ErrServerError)
		return
	}

	r.JSON(http.StatusOK, map[string]interface{}{"status": true, "data": "ok"})
}

func statsInfo(req *http.Request, r render.Render) {
	if req.Method != "GET" {
		r.JSON(http.StatusMethodNotAllowed, ErrBadRequest)
		return
	}
	stats := falcon.GetCurrentStatsDB()
	r.JSON(http.StatusOK, map[string]interface{}{"status": true, "data": stats})

}
