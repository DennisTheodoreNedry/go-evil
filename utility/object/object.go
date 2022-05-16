package object

import (
	"os"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func GetName() string {
	path, err := os.Executable()

	if err != nil {
		notify.Error(err.Error(), "object.GetName()")
	}

	program_name := ""
	for _, c := range path {
		if string(c) == "/" {
			program_name = ""
		} else {
			program_name += string(c)
		}
	}

	return program_name
}
