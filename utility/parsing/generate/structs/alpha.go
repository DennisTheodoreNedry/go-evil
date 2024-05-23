package structs

import (
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/functions"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
	"github.com/DennisTheodoreNedry/go-evil/utility/structure/structs"
)

// Generates the alpha struct (read rib) of the malware
func Generate_alpha(data_object *json.Json_t) {

	data_object.Add_go_struct(
		structs.Go_struct_t{
			Name:     "alpha_t",
			Contents: []string{"alphabet []string"},
		})

	body := []string{
		"to_return := \"\"",
		"for _, number := range value{",
		"to_return += obj.alphabet[number]",
		"}",
		"return to_return",
	}

	data_object.Add_go_function(functions.Go_func_t{Name: "construct_string", Func_type: "", Part_of_struct: "alpha_t",
		Return_type: "string", Parameters: []string{"value []int"}, Gut: body})

}
