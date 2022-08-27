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
	new_json.Malware_path = "./output/"
	new_json.Malware_gut = append(new_json.Malware_gut, "package main")
	return new_json
}

//
//
// Used when you want to send the data structure
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
// Used to convert the received data into workable data
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
