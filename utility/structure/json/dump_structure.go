package json

import (
	"encoding/json"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Creates byte code from our json structure
func (object *Json_t) Dump() []byte {
	serial_json, err := json.Marshal(object)

	if err != nil {
		notify.Error(err.Error(), "json.Convert_to_json()")
	}

	return serial_json
}
