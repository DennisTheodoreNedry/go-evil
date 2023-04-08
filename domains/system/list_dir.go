package system

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Reads the contents of a directory and places the result into a runtime variable
func list_dir(s_json string, value string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "list_dir"
	arr := structure.Create_evil_object(value)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(config []string){", function_call),
		"if len(config) < 2{",
		"spine.log(\"The provided evil array does not contain all required values\")",
		"return",
		"}",
		"obj_type := spine.variable.get(config[0])",
		"path := spine.variable.get(config[1])",
		"result, err := ioutil.ReadDir(path)",
		"if err == nil{",
		"evil_array := \"${\"",
		"for _, file := range result{",
		"if obj_type == \"file\" && !file.IsDir() || obj_type == \"dir\" && file.IsDir() || obj_type == \"\" {",
		"evil_array += fmt.Sprintf(\"\\\"%s/%s\\\",\", path, file.Name())",
		"}",
		"}",
		"evil_array += \"}$\"",
		"spine.variable.set(evil_array)",
		"}}"})

	data_object.Add_go_import("io/ioutil")
	data_object.Add_go_import("fmt")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	return []string{fmt.Sprintf("%s(%s)", function_call, arr.To_string("array"))}, structure.Send(data_object)

}
