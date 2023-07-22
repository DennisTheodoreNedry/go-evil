package json

import (
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
)

// Adds a go based function to the final go code
func (object *Json_t) Add_go_function(new_func functions.Go_func_t) {

	for _, calls := range object.GO_functions {
		if (calls.Name == new_func.Name) && (calls.Part_of_struct == new_func.Part_of_struct) {
			return
		}
	}

	object.GO_functions = append(object.GO_functions, new_func)
}
