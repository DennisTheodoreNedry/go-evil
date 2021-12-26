package encryption

import (
	"fmt"
	"os"

	"github.com/s9rA16Bf4/go-evil/utility/notify"
)

// We target files, folders to encrypt

type target_t struct {
	target_name       string
	encryption_method string
}

var c_target target_t

func Encrypt_set_target(path string) { // Either a file or a folder
	target_info, err := os.Stat(path)
	if err != nil {
		notify.Notify_error(err.Error(), "encrypt.Encrypt_set_target()")
	}
	fmt.Println(target_info)
}
