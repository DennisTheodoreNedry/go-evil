package debugger

import (
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

// Generate behavior function for debugging
func Generate_behavior(data_object *json.Json_t) {

	if data_object.Debugger_behavior != "none" {
		identify(data_object) // Adds necessary code to identify a debugger

		switch data_object.Debugger_behavior {
		case "stop":
			stop_behavior(data_object)
		case "remove":
			remove_behavior(data_object)
		case "loop":
			loop_behavior(data_object)
		}
	}

}
