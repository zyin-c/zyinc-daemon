package logs

import (
	"fmt"
	"log"
	"os"

	"github.com/coreos/go-systemd/v22/journal"
)

func init() {
}

func sendlogtojournal(message string, priority journal.Priority, identifier string) error {
	return journal.Send(message, priority, map[string]string{
		"SYSLOG_IDENTIFIER": identifier,
		"PROCESS_NAME":      identifier,
		"PROCESS_PID":       fmt.Sprintf("%d", os.Getpid()),
	})
}

func JournalDebug(identifier, message string) {
	if err := sendlogtojournal(fmt.Sprintf("[DEBUG] %s", message), journal.PriDebug, identifier); err != nil {
		log.Println(err)
	}
}

func JournalInfo(identifier, message string) {
	if err := sendlogtojournal(fmt.Sprintf("[INFO] %s", message), journal.PriInfo, identifier); err != nil {
		log.Println(err)
	}
}

func JournalWarn(identifier, message string) {
	if err := sendlogtojournal(fmt.Sprintf("[WARN] %s", message), journal.PriWarning, identifier); err != nil {
		log.Println(err)
	}
}

func JournalErr(identifier, message string) {
	if err := sendlogtojournal(fmt.Sprintf("[ERROR] %s", message), journal.PriErr, identifier); err != nil {
		log.Println(err)
	}
}

func JournalCritical(identifier, message string) {
	if err := sendlogtojournal(fmt.Sprintf("[CRITICAL] %s", message), journal.PriErr, identifier); err != nil {
		log.Println(err)
	}
}
