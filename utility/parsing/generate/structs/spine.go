package structs

import (
	"github.com/TeamPhoneix/go-evil/utility/structure/functions"
	"github.com/TeamPhoneix/go-evil/utility/structure/json"
	"github.com/TeamPhoneix/go-evil/utility/structure/structs"
)

// Generates the core struct (read spine) of each malware
func Generate_spine(data_object *json.Json_t) {

	data_object.Add_go_struct(structs.Go_struct_t{
		Name: "spine_t",
		Contents: []string{
			"variable  var_t",
			"crypt crypt_t",
			"path string",
			"alpha alpha_t",
			"logging_lvl string",
			"is_admin bool",
			"terminate bool",
			"return_code int",
		},
	})

	body := []string{}

	if data_object.Target_os == "windows" {
		body = append(body, "_, err := os.Open(\"\\\\.\\\\PHYSICALDRIVE0\")")
	} else {
		body = append(body, "_, err := os.Open(\"/etc/sudoers\")")
	}

	body = append(body, "if err != nil{", "obj.is_admin = false", "}else{", "obj.is_admin = true", "}")

	data_object.Add_go_function(functions.Go_func_t{Name: "check_privileges", Func_type: "", Part_of_struct: "spine_t",
		Return_type: "", Parameters: []string{}, Gut: body})

	data_object.Add_go_function(functions.Go_func_t{Name: "log", Func_type: "", Part_of_struct: "spine_t",
		Return_type: "", Parameters: []string{"msg string"}, Gut: []string{"notify.Log(msg, spine.logging_lvl, \"3\")"}})

	data_object.Add_go_import("os")
	data_object.Add_go_global("var spine spine_t")

}
