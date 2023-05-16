package parsing

import (
	"github.com/s9rA16Bf4/go-evil/utility/parsing/debugger"
	evil_final "github.com/s9rA16Bf4/go-evil/utility/parsing/finalize"
	evil_generate "github.com/s9rA16Bf4/go-evil/utility/parsing/generate"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Parses the contents of the provided file
func Parse(data_object *json.Json_t) {
	preface(data_object) // Handles every preface we could possibly want done before we start parsing

	evil_generate.Generate_structs(data_object)

	debugger.Generate_behavior(data_object)

	evil_generate.Generate_go_functions(data_object)

	evil_generate.Generate_main(data_object)

	evil_final.Construct_malware(data_object)

}
