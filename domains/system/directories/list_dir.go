package directories

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure"
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Reads the contents of a directory and places the result into a runtime variable
func List(value string, data_object *json.Json_t) []string {
	function_call := "list_dir"
	arr := structure.Create_evil_object(value)

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"config []string"},

		Gut: []string{
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
			"}",
		}})

	data_object.Add_go_import("io/ioutil")
	data_object.Add_go_import("fmt")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	return []string{fmt.Sprintf("%s(%s)", function_call, arr.To_string("array"))}

}
