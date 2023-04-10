package parsing

import (
	"github.com/TeamPhoneix/go-evil/utility/parsing/debugger"
	evil_final "github.com/TeamPhoneix/go-evil/utility/parsing/finalize"
	evil_generate "github.com/TeamPhoneix/go-evil/utility/parsing/generate"
)

// Parses the contents of the provided file
func Parse(s_json string) string {
	s_json = preface(s_json) // Handles every preface we could possibly want done before we start parsing

	s_json = evil_generate.Generate_structs(s_json)

	s_json = debugger.Generate_behavior(s_json)

	s_json = evil_generate.Generate_go_functions(s_json)

	s_json = evil_generate.Generate_main(s_json)

	s_json = evil_final.Construct_malware(s_json)

	return s_json
}
