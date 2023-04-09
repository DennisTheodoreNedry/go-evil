package debugger

import (
	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Generate behavior function for debugging
func Generate_behavior(s_json string) string {
	data_object := structure.Receive(s_json)

	if data_object.Debugger_behavior != "none" {
		s_json = identify(s_json) // Adds neccessary code to identify a debugger

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
