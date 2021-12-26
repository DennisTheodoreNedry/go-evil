package domains

import (
	"fmt"
	"os"

	"github.com/s9rA16Bf4/go-evil/utility/converter"
)

func System_exit(status_lvl string) {
	value := converter.String_to_int(status_lvl, "system.System_exit()")
	os.Exit(value)
}

func System_out(msg string) {
	fmt.Println(msg)
}
