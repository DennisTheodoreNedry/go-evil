package generate

import "github.com/TeamPhoneix/go-evil/utility/parsing/generate/structs"

// Generate different structs in the final malware
func Generate_structs(s_json string) string {
	s_json = structs.Generate_runtime_variable(s_json)
	s_json = structs.Generate_crypt(s_json)
	s_json = structs.Generate_alpha(s_json)
	s_json = structs.Generate_spine(s_json)

	return s_json
}
