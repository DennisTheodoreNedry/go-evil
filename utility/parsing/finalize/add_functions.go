package finalize

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

func Add_functions(s_json string) string {
	data_object := structure.Receive(s_json)

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
				header += fmt.Sprintf("%s,", parameter)
			}
		}
		header += ")"

		// Add the return type
		header += fmt.Sprintf(" %s ", new_func.Return_type)

		// Construct the body
		body = append(body, header, "{")
		body = append(body, new_func.Gut...)
		body = append(body, "}")

		data_object.Add_malware_lines(body)
	}

	return structure.Send(data_object)
}
