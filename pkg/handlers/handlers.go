package handlers

import (
	"log"

	"github.com/zyin-c/zyinc-daemon/pkg/runner"
)

var cmdRunner *runner.Runner
var err error

func init() {
	cmdRunner, err = runner.NewRunner()
	if err != nil {
		log.Fatalln(err)
	}
}
