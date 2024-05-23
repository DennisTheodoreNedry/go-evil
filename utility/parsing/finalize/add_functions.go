package finalize

import (
	"fmt"

	"github.com/DennisTheodoreNedry/go-evil/utility/structure/json"
)

func Add_functions(data_object *json.Json_t) {

	for _, new_func := range data_object.GO_functions {
		body := []string{}

		// Construct header
		header := "func "
		if new_func.Part_of_struct != "" { // This function is part of a structure
			header += fmt.Sprintf(" (obj *%s)", new_func.Part_of_struct)
		}

		header += fmt.Sprintf(" %s(", new_func.Name)

		// Parameters
		if len(new_func.Parameters) > 0 {
			for _, parameter := range new_func.Parameters {
				if parameter == "" { // It's empty so skip this one
					continue
				}

				header += fmt.Sprintf("%s,", parameter)
			}
		}
		header += ")"

		// Add the return type
		header += fmt.Sprintf(" %s {", new_func.Return_type)

		// Construct the body
		body = append(body, header)
		body = append(body, new_func.Gut...)

		switch new_func.Return_type {
		case "int":
			body = append(body, "return 0")

		case "bool":
			body = append(body, "return true")

		case "string":
			body = append(body, "return \"\"")

		}

		body = append(body, "}")

		data_object.Add_malware_lines(body)
	}

}
