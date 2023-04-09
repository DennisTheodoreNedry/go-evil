package debugger

import "github.com/TeamPhoneix/go-evil/utility/structure"

// Generate the debugger detection function
func stop_behavior(s_json string) string {
	data_object := structure.Receive(s_json)
	body := []string{"func stop_behavior() bool {",
		"toReturn := false",
		"toReturn = detect_debugger()",
		"if toReturn {",
		"os.Exit(42)",
		"}"}

	body = append(body, "return toReturn", "}")
	data_object.Add_go_function(body)
	data_object.Add_go_import("os")

	return structure.Send(data_object)
}
