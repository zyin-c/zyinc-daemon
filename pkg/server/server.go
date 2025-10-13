package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/zyin-c/extras/schema"
	"github.com/zyin-c/zyinc-daemon/common/config"
)

func NewServer(cfg ServerConfig) *Server {
	if strings.TrimSpace(cfg.HTTPHost) == "" {
		cfg.HTTPHost = config.DEFAULT_HOST
	}

	if strings.TrimSpace(cfg.SocketPath) == "" {
		cfg.SocketPath = config.DEFAULT_SOCKETPATH
	}
	if cfg.MaxWorkers <= 0 {
		cfg.MaxWorkers = config.DEFAULT_WORKERCOUNT
	}

	return &Server{
		socketPath:    cfg.SocketPath,
		httpHost:      cfg.HTTPHost,
		useHttp:       cfg.UseHTTP,
		handler:       make(map[string]EventHandler),
		maxWorkers:    cfg.MaxWorkers,
		connSemaphore: make(chan struct{}, cfg.MaxWorkers),
	}
}

func (s *Server) On(event string, handler EventHandler) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handler[event] = handler
}

func (s *Server) handleEvent(w io.Writer, payload schema.SocketMessage) {
	s.mu.RLock()
	handler, ok := s.handler[payload.Event]
	s.mu.RUnlock()

	if !ok {
		fmt.Fprintf(w, `{"error": "no handler for evente %q"}`+"\n", payload.Event)
		return
	}

	handler(w, payload)
}

func (s *Server) handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		<-s.connSemaphore
		s.wg.Done()
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		var payload schema.SocketMessage
		line := scanner.Bytes()
		if err := json.Unmarshal(line, &payload); err != nil {
			fmt.Fprintf(conn, `{"error": "invalid JSON: %s"}`+"\n", err.Error())
			continue
		}
		go s.handleEvent(conn, payload)
	}
}

func (s *Server) serveHttp(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) Start() error {
	if s.useHttp {
		return s.startHttp()
	}
	return s.startUnixSocketServer()
}

func (s *Server) startUnixSocketServer() error {
	if _, err := os.Stat(s.socketPath); err == nil {
		os.Remove(s.socketPath)
	}

	listener, err := net.Listen("unix", s.socketPath)
	if err != nil {
		return fmt.Errorf("failed to listen on socket: %w", err)
	}
	s.listener = listener
	fmt.Println("Unix socket server listening on:", s.socketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) startHttp() error {
	http.HandleFunc("/", s.serveHttp)
	fmt.Println("HTTP server listening on:", s.httpHost)
	return http.ListenAndServe(s.httpHost, nil)
}
