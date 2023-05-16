package systemcommands

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Reboots the computer
func Reboot(data_object *json.Json_t) []string {
	function_call := "Reboot"

	cmd := ""
	if data_object.Target_os == "windows" {
		cmd = "shutdown /r"
	} else {
		cmd = "shutdown -r now"
	}

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{},
		Gut: []string{
			fmt.Sprintf("exec.Command(\"%s\").Run()", cmd),
		}})

	data_object.Add_go_import("os/exec")

	return []string{fmt.Sprintf("%s()", function_call)}

}
