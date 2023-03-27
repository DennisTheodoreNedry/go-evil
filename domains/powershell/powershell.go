package powershell

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func set_execution_policy(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	system_call := "set_execution_policy"
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

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int){", system_call),
		"policy := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"err := exec.Command(\"powershell\", \"Set-ExecutionPolicy\", policy).Run()",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"}"})

	data_object.Add_go_import("os/exec")
	parameter_1 := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", system_call, parameter_1)}, structure.Send(data_object)
}
