package handlers

import (
	"fmt"
	"io"

	"github.com/zyin-c/extras/schema"
	"github.com/zyin-c/zyinc-daemon/common/utils"
)

func PsEventHandler(w io.Writer, data schema.SocketMessage) {

	retstring, err := utils.ReturnData(schema.SocketResponse{})
	if err != nil {
		fmt.Fprintf(w, `{"error": %v}`+"\n", err)
	}
	fmt.Fprint(w, retstring+"\n")
}
