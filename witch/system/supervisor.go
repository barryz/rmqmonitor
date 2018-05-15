package system

import (
	"strings"
)

// Supervisor is the supervisor process control system.
type Supervisor struct {
	name    string
	service string
}

// NewSupervisor creates new Supervisor instance.
func NewSupervisor(service string) *Supervisor {
	return &Supervisor{
		name:    "/usr/bin/supervisorctl",
		service: service,
	}
}

// IsAlive gets results from `supervisorctl status [service]`
func (s *Supervisor) IsAlive() (int, bool) {
	r, err := ExecCommand(s.name, []string{"status", s.service})
	if err != nil {
		return -1, false
	}
	return -1, strings.Contains(r, "RUNNING")
}

// Start executes `supervisorctl start [service]`
func (s *Supervisor) Start() (bool, error) {
	_, err := ExecCommand(s.name, []string{"start", s.service})
	return err == nil, err
}

// Restart executes `supervisorctl restart [service]`
func (s *Supervisor) Restart() (bool, error) {
	_, err := ExecCommand(s.name, []string{"restart", s.service})
	return err == nil, err
}

// Stop executes `supervisorctl stop [service]`
func (s *Supervisor) Stop() bool {
	_, err := ExecCommand(s.name, []string{"stop", s.service})
	return err == nil
}
