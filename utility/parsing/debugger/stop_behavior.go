package debugger

import (
	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
)

// Generate the debugger detection function
func stop_behavior(s_json string) string {
	data_object := structure.Receive(s_json)

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

	return structure.Send(data_object)
}
