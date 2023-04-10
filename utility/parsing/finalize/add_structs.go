package finalize

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

func Add_structs(s_json string) string {
	data_object := structure.Receive(s_json)

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

	return structure.Send(data_object)
}
