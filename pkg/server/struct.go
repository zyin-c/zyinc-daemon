package server

import (
	"io"
	"net"
	"sync"

	"github.com/zyin-c/extras/schema"
)

type EventHandler func(w io.Writer, data schema.SocketMessage)

type Server struct {
	socketPath    string
	httpHost      string
	useHttp       bool
	listener      net.Listener
	handler       map[string]EventHandler
	mu            sync.RWMutex
	wg            sync.WaitGroup
	maxWorkers    int
	connSemaphore chan struct{}
}

type ServerConfig struct {
	SocketPath string
	HTTPHost   string
	UseHTTP    bool
	MaxWorkers int //default to 10
}
