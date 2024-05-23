package json

import (
	"encoding/json"

	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Creates byte code from our json structure
func (object *Json_t) Dump() []byte {
	serial_json, err := json.Marshal(object)

	if err != nil {
		notify.Error(err.Error(), "json.Convert_to_json()", 1)
	}

	return serial_json
}
