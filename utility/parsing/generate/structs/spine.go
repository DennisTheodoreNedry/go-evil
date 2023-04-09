package structs

import "github.com/TeamPhoneix/go-evil/utility/structure"

// Generates the core struct (read spine) of each malware
func Generate_spine(s_json string) string {
	data_object := structure.Receive(s_json)

	data_object.Add_go_struct([]string{
		"type spine_t struct {",
		"variable  var_t",
		"crypt crypt_t",
		"path string",
		"alpha alpha_t",
		"logging_lvl string",
		"is_admin bool",
		"}"})

	body := []string{"func (obj *spine_t) check_privileges(){"}

	if data_object.Target_os == "windows" {
		body = append(body, "_, err := os.Open(\"\\\\.\\\\PHYSICALDRIVE0\")")
	} else {
		body = append(body, "_, err := os.Open(\"/etc/sudoers\")")
	}

	body = append(body, "if err != nil{", "obj.is_admin = false", "}else{", "obj.is_admin = true", "}", "}")

	body = append(body, "func (obj *spine_t) log(msg string){", "notify.Log(msg, spine.logging_lvl, \"3\")", "}")

	data_object.Add_go_function(body)
	data_object.Add_go_import("os")
	data_object.Add_go_global("var spine spine_t")
	return structure.Send(data_object)

}
