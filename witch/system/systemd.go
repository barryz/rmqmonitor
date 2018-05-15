package system

import (
	"strings"
)

// Systemd is the systemd process control system.
type Systemd struct {
	name    string
	service string
}

// NewSystemd creates new Systemd instance.
func NewSystemd(service string) *Systemd {
	return &Systemd{
		name:    "/bin/systemctl",
		service: service,
	}
}

// IsAlive gets results from `systemctl is-active [service]`
func (s *Systemd) IsAlive() (int, bool) {
	r, err := ExecCommand(s.name, []string{"is-active", s.service})
	if err != nil {
		return -1, false
	}
	return -1, strings.HasPrefix(r, "active")
}

// Start executes `systemctl start [service]`
func (s *Systemd) Start() (bool, error) {
	_, err := ExecCommand(s.name, []string{"start", s.service})
	return err == nil, err
}

// Restart executes `systemctl restart [service]`
func (s *Systemd) Restart() (bool, error) {
	_, err := ExecCommand(s.name, []string{"restart", s.service})
	return err == nil, err
}

// Stop executes `systemctl stop [service]`
func (s *Systemd) Stop() bool {
	_, err := ExecCommand(s.name, []string{"stop", s.service})
	return err == nil
}
