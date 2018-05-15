package system

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
	"time"
)

// Launcher supervises the process status, start, stop and restart.
type Launcher struct {
	pidFile string
	cmd     string
}

// NewLauncher creates new system.
func NewLauncher(pidFile, cmd string) *Launcher {
	return &Launcher{
		pidFile: pidFile,
		cmd:     cmd,
	}
}

func (s *Launcher) writePid(pid int) {
	if err := WriteFile(s.pidFile, []byte(strconv.FormatInt(int64(pid), 10)), 0644); err != nil {
		log.Fatalf("Failed to write pid file: %s", err)
	}
}

func (s *Launcher) readPid() (int, bool) {
	f, err := ioutil.ReadFile(s.pidFile)
	if err != nil {
		log.Printf("Error reading pid file[%s]: %s", s.pidFile, err)
		return -1, false
	}

	pid, err := strconv.Atoi(string(f))
	if err != nil {
		log.Printf("Invalid pid value[%s]: %s", s.pidFile, err)
		return -1, false
	}

	return pid, true
}

func (s *Launcher) pidAlive(pid int) bool {
	return syscall.Kill(pid, 0) == nil
}

// IsAlive check if the process alive.
func (s *Launcher) IsAlive() (int, bool) {
	pid, ok := s.readPid()
	if !ok || pid < 1 {
		return pid, false
	}
	return pid, s.pidAlive(pid)
}

// Start starts the process.
func (s *Launcher) Start() (bool, error) {
	if pid, ok := s.IsAlive(); ok {
		log.Printf("The process is alive, pid: %d", pid)
		return true, nil
	}

	log.Printf("Starting [%s]", s.cmd)
	child := exec.Command("/bin/bash", []string{"-c", s.cmd}...)
	child.Stdin = os.Stdin
	child.Stdout = os.Stdout
	child.Stderr = os.Stderr
	if err := child.Start(); err != nil {
		log.Printf("Failed to start: %s", err)
		return false, err
	}
	s.writePid(child.Process.Pid)
	go child.Wait()
	return true, nil
}

// Restart restart the process
func (s *Launcher) Restart() (bool, error) {
	s.Stop()
	return s.Start()
}

// Stop stops the process.
func (s *Launcher) Stop() bool {
	pid, ok := s.IsAlive()
	if !ok {
		log.Printf("The process not alive")
		return true
	}
	syscall.Kill(pid, syscall.SIGTERM)
	stopped := make(chan bool)
	go func() {
		for s.pidAlive(pid) {
			time.Sleep(time.Second)
		}
		close(stopped)
	}()
	select {
	case <-stopped:
		log.Printf("[INFO] Stop the process success.")
	case <-time.After(time.Duration(stopWaitSecs) * time.Second):
		log.Printf("[INFO] Stop the process timeout, force to kill.")
		syscall.Kill(pid, syscall.SIGKILL)
	}
	return true
}

// WriteFile tries to create parent directory before WriteFile.
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	if err := os.MkdirAll(path.Dir(filename), 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, perm)
}
