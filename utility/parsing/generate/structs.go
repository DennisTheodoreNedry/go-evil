package generate

import (
	"github.com/s9rA16Bf4/go-evil/utility/parsing/generate/structs"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Generate different structs in the final malware
func Generate_structs(data_object *json.Json_t) {

	structs.Generate_runtime_variable(data_object)
	structs.Generate_crypt(data_object)
	structs.Generate_alpha(data_object)
	structs.Generate_spine(data_object)

}
