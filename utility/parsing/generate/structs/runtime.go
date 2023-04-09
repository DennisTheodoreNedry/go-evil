package structs

import "github.com/TeamPhoneix/go-evil/utility/structure"

// Generates the runtime variable struct (read rib) of the malware
func Generate_runtime_variable(s_json string) string {
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
		"for _, value := range result {",
		"i_number := tools.String_to_int(value[2])",
		"grabbed_value := \"NULL\"",
		"if i_number != -1 {",
		"if i_number > 0 && i_number < 5 {",
		"grabbed_value = obj.values[i_number-1]",
		"} else {",
		"switch i_number {",

		"case 13:",
		"grabbed_value = obj.foreach",

		"case 23:",
		"grabbed_value = tools.Grab_executable_name()",

		"case 39:",
		"grabbed_value = tools.Grab_CWD()",

		"case 40:",
		"grabbed_value = tools.Grab_home_dir()",

		"case 666:",
		"grabbed_value = tools.Grab_username()",

		"default:",
		"spine.log(fmt.Sprintf(\"Error, unknown value '%d'\", i_number))",

		"}",
		"}",
		"line = strings.ReplaceAll(line, value[1], grabbed_value)",
		"}",
		"}",
		"toReturn = line",
		"}",
		"return toReturn",
		"}"})

	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")
	data_object.Add_go_import("regexp")
	data_object.Add_go_import("strings")
	data_object.Add_go_import("fmt")

	data_object.Add_go_const("GRAB_VAR = \"(€([0-9]+)€)\"")

	return structure.Send(data_object)
}
