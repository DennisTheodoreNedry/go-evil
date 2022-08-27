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

	// Add boot functions
	for _, boot_name := range boot_functions {
		main_functions = append(main_functions, fmt.Sprintf("%s()", boot_name))
	}

	main_functions = append(main_functions, "for {")

	// Add loop function
	for _, loop_name := range loop_functions {
		main_functions = append(main_functions, (fmt.Sprintf("%s()", loop_name)))
	}
	main_functions = append(main_functions, "}", "}")

	data_object.Add_go_function(main_functions)

	return structure.Send(data_object)

}

//
//
// Identify sub function type and populates two string arrays
//
//
func identify_sub_function(d_funcs []structure.Func_t) ([]string, []string) {
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

	boot_func, loop_func := identify_sub_function(data_object.Functions)

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
// Parses the contents of the provided file
//
func Parse(s_json string) string {
	s_json = preface(s_json) // Handles every preface we could possibly want done before we start parsing

	s_json, boot_func, loop_func := generate_sub_functions(s_json)

	s_json = generate_main_function(s_json, boot_func, loop_func)

	data_object := structure.Receive(s_json)
	s_json = fill_malware_gut(structure.Send(data_object))

	return s_json
}
