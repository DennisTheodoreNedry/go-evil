package version

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EVIL_VERSION = "2.0.0"
	TEXT_EDITOR  = "1.0.0"
)

func Version() {
	notify.Inform(fmt.Sprintf("Compiler version: %s", EVIL_VERSION))
	notify.Inform(fmt.Sprintf("Text editor version: %s", TEXT_EDITOR))
}
