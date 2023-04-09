package structs

import "github.com/TeamPhoneix/go-evil/utility/structure"

// Generates the alpha struct (read rib) of the malware
func Generate_alpha(s_json string) string {
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
