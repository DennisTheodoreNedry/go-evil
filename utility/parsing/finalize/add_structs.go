package finalize

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

func Add_structs(data_object *json.Json_t) {

	for _, new_struct := range data_object.GO_struct {
		body := []string{}

		// Construct header
		header := fmt.Sprintf("type %s struct {", new_struct.Name)

		// Construct the body
		body = append(body, header)
		body = append(body, new_struct.Contents...)
		body = append(body, "}")

		data_object.Add_malware_lines(body)
	}

}
