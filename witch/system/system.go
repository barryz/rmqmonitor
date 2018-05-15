package system

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

var (
	stopWaitSecs = 5
)

// System is the interface of process control system.
type System interface {
	// IsAlive checks process is alive.
	IsAlive() (int, bool)
	// Start starts process.
	Start() (bool, error)
	// Start restart process.
	Restart() (bool, error)
	// Stop stops process.
	Stop() bool
}

type Stats interface {
	// Reset the RabbitMQ StatsDB, A node may be randomly selected
	Reset() (bool, string, error)
	// Terminate the RabbitMQ StatsDB, Not to choose the node
	Terminate() (bool, string, error)
	// Crash the RabbitMQ StatsDB, Not to choose the node
	Crash() (bool, string, error)
}

// Action is the system action.
type Action struct {
	Name string `json:"name"`
}

// ActionStatus is the status of action.
type ActionStatus struct {
	Status bool   `json:"status"`
	Text   string `json:"text"`
}

// SysController controls the System.
type SysController struct {
	System
}

// StatsController controls the RabbitMQ StatsDB
type StatsController struct {
	Stats
}

func (c *StatsController) Handle(action *Action) *ActionStatus {
	var (
		st     = &ActionStatus{}
		err    error
		output string
	)

	switch action.Name {
	case "status":
		fallthrough
	case "reset":
		if st.Status, output, err = c.Reset(); err != nil {
			st.Text = err.Error()
		} else {
			st.Text = output
		}
	case "terminate":
		if st.Status, output, err = c.Terminate(); err != nil {
			st.Text = err.Error()
		} else {
			st.Text = output
		}
	case "crash":
		if st.Status, output, err = c.Crash(); err != nil {
			st.Text = err.Error()
		} else {
			st.Text = output
		}

	default:
		st.Status, st.Text = false, fmt.Sprintf("Invalid action: %s", action.Name)

	}
	log.Println("[INFO]: StatsDB Action finished")
	return st
}

// Handle plays action.
func (c *SysController) Handle(action *Action) *ActionStatus {
	var (
		st  = &ActionStatus{}
		err error
	)
	switch action.Name {
	case "status":
		fallthrough
	case "is_alive":
		_, st.Status = c.IsAlive()
	case "start":
		if st.Status, err = c.Start(); err != nil {
			st.Text = err.Error()
		}
	case "stop":
		st.Status = c.Stop()
	case "restart":
		if st.Status, err = c.Restart(); err != nil {
			st.Text = err.Error()
		}
	default:
		st.Status, st.Text = false, fmt.Sprintf("Invalid action: %s", action.Name)
	}
	log.Printf("[INFO] System Action finished")
	return st
}

func ExecCommand(name string, args []string) (string, error) {
	var buf bytes.Buffer
	log.Printf("[INFO] Exec %s %v", name, args)
	child := exec.Command(name, args...)
	child.Stdout = &buf
	child.Stderr = &buf
	if err := child.Start(); err != nil {
		log.Printf("[ERROR] Failed to start: %s", err)
		return buf.String(), err
	}
	child.Wait()
	return buf.String(), nil
}
