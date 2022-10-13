package structure

import (
	"encoding/base64"
	"encoding/json"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// Creates a json object and returns it
//
//
func Create_json_object() json_t {
	var new_json json_t
	new_json.Malware_src_file = "malware.go"
	new_json.Malware_gut = append(new_json.Malware_gut, "package main")

	// Settings for all the webview instances
	new_json.Width = 800
	new_json.Height = 600

	// Used when we are doing a 1:1 mapping of our js function to a evil one
	new_json.Bind_gut = make(map[string]string)

	new_json.Var_max = 5

	for id := 0; id < new_json.Var_max; id++ {
		var new_var Compile_var_t
		new_var.Set_value("NULL") // Default value
		new_json.Comp_var = append(new_json.Comp_var, new_var)
	}

	return new_json
}

//
//
// Creates an evil array object and returns it
//
//
func Create_evil_object(arr_content string) Evil_array_t {
	var new_arr Evil_array_t
	new_arr.length = 0
	new_arr.gut = []string{}

	new_arr.Parse(arr_content) // Will populate the array with the provided gut

	return new_arr
}

//
//
// Serializes the json structure into a base64 string which is ready to be sent
//
//
func Send(object json_t) string {
	serial_json, err := json.Marshal(object)

	if err != nil {
		notify.Error(err.Error(), "json.Convert_to_json()")
	}

	return base64.StdEncoding.EncodeToString(serial_json)
}

//
//
// Used to convert the received serialized json structure into workable data
//
//
func Receive(object string) json_t {
	serialize_json, err := base64.StdEncoding.DecodeString(object)

	if err != nil {
		notify.Error(err.Error(), "json.Receive()")
	}

	var result json_t
	if err := json.Unmarshal(serialize_json, &result); err != nil {
		notify.Error(err.Error(), "json.Convert_to_json_t()")
	}

	return result
}
