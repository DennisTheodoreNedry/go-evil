package converter

import (
	"strconv"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func String_to_int(val string, c_func string) int {
	value, err := strconv.Atoi(val) // Can we convert it
	if err != nil {
		notify.Error("Failed to convert "+val+" to integer", c_func)
	}
	notify.Log("Converted '"+val+"' to an integer", notify.Verbose_lvl, "3")
	return value
}
