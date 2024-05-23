package version

import (
	"fmt"
	"os"

	notify "github.com/DennisTheodoreNedry/notify_handler"
)

const (
	EVIL_VERSION       = "3.0.1"
	EVIL_VERSION_SMALL = "3.0"
)

func Version(value string) string {
	notify.Inform(fmt.Sprintf("Compiler version: %s", EVIL_VERSION))
	os.Exit(0)
	return ""
}
