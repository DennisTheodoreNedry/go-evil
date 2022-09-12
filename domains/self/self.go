package self

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

//
//
// Calls on a function declared
//
//
func Call_function(func_name string, s_json string) (string, string) {
	data_object := structure.Receive(s_json)
	call := "call()"

	if data_object.Obfuscate {
		call = fmt.Sprintf("%s()", tools.Generate_random_string())

		// We will also need to find the newly named function
		for _, def_func := range data_object.Functions {

			if def_func.Func_type == "c" {
				for _, prev_name := range def_func.Prev_names {
					if prev_name == func_name {
						func_name = def_func.Get_name()
						break
					}
				}
			}

		}

	}

	func_name = tools.Erase_delimiter(func_name, `"`) // Removes all " from the string

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", call),
		fmt.Sprintf("%s()", func_name),
		"}",
	})

	return call, structure.Send(data_object)
}
