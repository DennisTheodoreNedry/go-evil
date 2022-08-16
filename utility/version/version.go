package version

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EVIL_VERSION = "0.0.0"
)

func Version() {
	notify.Inform(fmt.Sprintf("Current version: %s", EVIL_VERSION))
}
