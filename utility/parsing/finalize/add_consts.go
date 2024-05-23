package finalize

import "github.com/DennisTheodoreNedry/go-evil/utility/structure/json"

func Add_consts(data_object *json.Json_t) {

	if len(data_object.GO_const) > 0 {
		data_object.Add_malware_line("const (")
		for _, new_const := range data_object.GO_const {
			data_object.Add_malware_line(new_const)
		}
		data_object.Add_malware_line(")")
	}

}
