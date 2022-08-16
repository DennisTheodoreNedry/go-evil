package system

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
// Exits the malware
//
func Exit(return_code string) {
	value := tools.String_to_int(return_code)
	os.Exit(value)
}

//
// Prints a message to the screen
//
func Out(msg string) {
	fmt.Println(msg)
}

//
// Executes a command on the OS and prints the result
//
func Exec(cmd string) {
	out, err := exec.Command(cmd).Output()

	if err != nil {
		notify.Warning(err.Error())
	} else {
		notify.Inform(string(out[:]))
	}
}
