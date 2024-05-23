package finalize

import "github.com/DennisTheodoreNedry/go-evil/utility/structure/json"

func Add_globals(data_object *json.Json_t) {

	if len(data_object.GO_global) > 0 {
		for _, new_global := range data_object.GO_global {
			data_object.Add_malware_line(new_global)
		}
	}

}
