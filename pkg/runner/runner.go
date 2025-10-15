package runner

import (
	"context"
	"os/exec"
	"sync"
	"time"
)

type ProcessState string

const (
	StateRunning ProcessState = "running"
	StateStopped ProcessState = "stopped"
	StateExited  ProcessState = "exited"
	StateFailed  ProcessState = "failed"
)

type RunnerInstance struct {
	BaseDirectory string `json:"basedir"`
	Filename      string `json:"filename"`
	EnvFile       string `json:"envfile"`
}

type ProcessInfo struct {
	Cmd       *exec.Cmd
	RunnerID  string
	State     ProcessState
	Instance  RunnerInstance
	StartTime time.Time
	StopTime  time.Time
	Pid       int
	ExitCode  int
	Restarts  int
	mu        sync.RWMutex
	cancel    context.CancelFunc
	ctx       context.Context
}

type Runner struct {
	Processes map[string]*ProcessInfo
	mu        sync.RWMutex
}

func NewRunner() (*Runner, error) {
	return &Runner{}, nil
}
