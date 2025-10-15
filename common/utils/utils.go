package utils

import (
	"encoding/json"

	"github.com/zyin-c/extras/schema"
)

// func SendToJournal(message string) {
// 	journal.Send(message,journal.Priority())
// }

func ReturnData(returndata schema.SocketResponse) (string, error) {
	bytes, err := json.Marshal(returndata)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
