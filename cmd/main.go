package main

import (
	"fmt"
	"io"
	"log"

	"github.com/zyin-c/extras/schema"
	"github.com/zyin-c/zyinc-daemon/pkg/server"
)

func main() {
	app := server.NewServer(server.ServerConfig{})

	app.On("ping", func(w io.Writer, payload schema.SocketMessage) {
		log.Println("request came")
		log.Println(payload)
		fmt.Fprintf(w, `{"pong": true, "event": "ping"}`+"\n")
	})

	if err := app.Start(); err != nil {
		panic(err)
	}

}
