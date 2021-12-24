package domains

import (
	"fmt"
	"os"
	"strconv"

	"github.com/s9rA16Bf4/go-evil/utility/notify"
)

func System_exit(status_lvl string) {
	value, err := strconv.Atoi(status_lvl)
	if err != nil {
		notify.Notify_error("Failed to convert "+status_lvl+" to integer", "parser.interpreter()")
	}
	os.Exit(value)
}

func System_out(msg string) {
	fmt.Println(msg)
}
