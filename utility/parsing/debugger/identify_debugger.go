package debugger

import (
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Adds the responsible code used for detecting if the malware is launched
// under a debugger
func identify(data_object *json.Json_t) {
	body := []string{"return coldfire.SandboxAll()"}

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_function(functions.Go_func_t{Name: "detect_debugger", Func_type: "null", Return_type: "bool", Gut: body})

}
