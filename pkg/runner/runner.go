package runner

import "github.com/zyin-c/extras/schema"

type Runner struct {
	RunnerId string
	Message  schema.SocketMessage
}

func NewRunner() (*Runner, error) {
	return &Runner{}, nil
}
