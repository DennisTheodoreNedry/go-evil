package json

import "github.com/TeamPhoneix/go-evil/utility/structure/functions"

// Adds a function to the structure
func (object *Json_t) Add_function(name string, f_type string, return_type string, gut []string) {
	var new_func functions.Func_t

	new_func.Set_name(name)

	new_func.Set_type(f_type)

	new_func.Set_return_type(return_type)

	new_func.Add_lines(gut)

	object.Functions = append(object.Functions, new_func)
}
