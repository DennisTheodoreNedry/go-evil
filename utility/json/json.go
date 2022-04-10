package json

import (
	"encoding/json"
	"time"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type data_t struct {
	Creation_Time string `json:"CurTime"` // Time of creation of the json structure
	Update_Time   string `json:"UpdTime"` // Time of latest update

	File        string `json:"File"`  // file path to the file we are reading
	Target_OS   string `json:"TOS"`   // The OS you're targeting
	Target_ARCH string `json:"TARCH"` // Target architecture
	Host_OS     string `json:"HOS"`   // The current running os

	Binary_name string `json:"BinName"` // The binary malware name
	Extension   string `json:"Ext"`     // Extension of the malware

	Debug    bool   `json:"Debug"`    // Debug mode
	TestMode string `json:"TestMode"` // Test mode, it's different from debug

	Verbose_LVL  string   `json:"VLVL"`        // Verbose level
	Call_history []string `json:"CallHistory"` // A string array containing all the functions we have passed through

	Malware_gut      []string `json:"MalwareLine"` // The content of the malware
	File_gut         []string `json:"FileLine"`    // Contents of the file we read in
	Imported_headers []string `json:"Headers"`     // The imported headers
}

func (json_object *data_t) Set_creation_time() {
	json_object.Creation_Time = time.Now().String()
}
func (json_object *data_t) Set_update_time() {
	json_object.Update_Time = time.Now().String()
}

func (json_object *data_t) Append_to_call(new_call string) {
	json_object.Call_history = append(json_object.Call_history, new_call)
}
func (json_object *data_t) Append_to_header(new_header string) {
	json_object.Imported_headers = append(json_object.Imported_headers, new_header)
}

func Create_object() data_t {
	var new_json data_t
	new_json.Set_creation_time() // Sets the time
	return new_json
}

func Convert_to_json(json_object data_t) []byte {
	serial_json, err := json.Marshal(json_object)
	if err != nil {
		notify.Error(err.Error(), "<call_function>")
		var null []byte
		return null
	}
	return serial_json
}

func Convert_to_data_t(in_json []byte) data_t {
	var result data_t
	err := json.Unmarshal(in_json, &result)
	if err != nil {
		notify.Error(err.Error(), "<call_function>")
		return data_t{}
	}
	return result
}
