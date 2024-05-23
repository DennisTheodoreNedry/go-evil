package debugger

import (
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Generates the code which will cause the malware to enter an infinite loop
func loop_behavior(data_object *json.Json_t) {
	body := []string{
		"toReturn := false",
		"toReturn = detect_debugger()",
		"if toReturn {",
		"for {",
		"}",
		"}",
		"return toReturn",
	}

	data_object.Add_go_function(functions.Go_func_t{Name: "loop_behavior", Func_type: "null", Return_type: "bool", Gut: body})

}
