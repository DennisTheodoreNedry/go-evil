package domains

import (
	"os"
)

func System_exit(status_lvl int) {
	os.Exit(status_lvl)
}
