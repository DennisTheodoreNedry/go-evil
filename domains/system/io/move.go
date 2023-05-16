package io

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure"
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Moves the target file to it's new location
func Move(value string, data_object *json.Json_t) []string {
	function_call := "move"

	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Obtained evil array had size %d, but 2 was requested", arr.Length()), "system.move()")
	}

	old_path := arr.Get(0)
	new_path := arr.Get(1)

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1  []int", "repr_2 []int"},
		Gut: []string{
			"old_path := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"new_path := spine.variable.get(spine.alpha.construct_string(repr_2))",

			"err := os.Rename(old_path, new_path)",

			"if err != nil{",
			"spine.log(err.Error())",
			"return",
			"}",
		}})

	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	// Construct our int array
	old_parameter := data_object.Generate_int_array_parameter(old_path)

	// Construct our int array
	new_parameter := data_object.Generate_int_array_parameter(new_path)

	return []string{fmt.Sprintf("%s(%s, %s)", function_call, old_parameter, new_parameter)}

}
