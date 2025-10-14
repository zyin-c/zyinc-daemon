package main

import (
	"log"

	"github.com/zyin-c/zyinc-daemon/pkg/route"
)

func main() {
	if err := route.StartServer(); err != nil {
		log.Fatalln("[ERR]", err)
	}
}
