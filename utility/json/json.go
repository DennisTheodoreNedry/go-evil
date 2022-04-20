package json

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type data_t struct {
	Creation_Time string `json:"CurTime"` // Time of creation of the json structure
	Update_Time   string `json:"UpdTime"` // Time of latest update

	File       string   `json:"File"`         // file path to the file we are reading
	File_gut   []string `json:"FileLine"`     // Contents of the file we read in
	File_lines int      `json:"FileRowCount"` // Amount of rows in the file
	File_count int      `json:"FileCount"`    // The amount of rows in `File_gut`

	Target_OS   string `json:"TOS"`   // The OS you're targeting
	Target_ARCH string `json:"TARCH"` // Target architecture
	Host_OS     string `json:"HOS"`   // The current running os

	Binary_name string `json:"BinName"` // The binary malware name
	Extension   string `json:"Ext"`     // Extension of the malware

	Debug    bool `json:"Debug"`    // Debug mode
	TestMode bool `json:"TestMode"` // Test mode, it's different from debug and can be seen as the counterpart to production mode

	Verbose_LVL  string   `json:"VerLV"`       // Verbose level
	Call_history []string `json:"CallHistory"` // A string array containing all the functions we have passed through

	Malware_gut []string `json:"MalwareLine"` // The content of the malware

	File_Domains     []string `json:"UsedDomains"`     // The domains that were used in the extracted file
	Imported_headers []string `json:"ImportedHeaders"` // The imported headers

	Compile_Time_variables []string `json:"CompTimeVar"`   // The variables themself
	Compile_Time_value     []string `json:"CompTimeValue"` // The values of the defined compile time variables
	Compile_Time_amount    int      `json:"CompTimeAmo"`   // The amount of defined compile time variables

	Run_Time_variables []string `json:"RunTimeVar"` // The variables themself
	Run_Time_amount    int      `json:"RunTimeAmo"` // The amount of defined run time variables
}

func (json_object *data_t) Set_creation_time() {
	json_object.Creation_Time = time.Now().String()
}
func (json_object *data_t) Set_update_time() {
	json_object.Update_Time = time.Now().String()
}

func (json_object *data_t) Append_to_call(new_call string) {
	json_object.Call_history = append(json_object.Call_history, new_call)
	json_object.Set_update_time()
}

func (json_object *data_t) Append_new_imported_domain(new_header string) {
	json_object.Imported_headers = append(json_object.Imported_headers, new_header)
	json_object.Set_update_time()
}

func (json_object *data_t) Append_file_gut(new_line string) {
	json_object.File_gut = append(json_object.File_gut, new_line)
	json_object.Set_update_time()
	json_object.Add_file_row()
}

func (json_object *data_t) Add_file_row() {
	json_object.File_lines += 1
	json_object.Set_update_time()
}

func (json_object *data_t) Append_File_domain(new_domain string) {
	json_object.File_Domains = append(json_object.File_Domains, new_domain)
	json_object.Set_update_time()
}

func (json_object *data_t) Get_File_domain() []string {
	return json_object.File_Domains
}

func (json_object *data_t) Append_malware_gut(new_line string) {
	json_object.Malware_gut = append(json_object.Malware_gut, new_line)
	json_object.Set_update_time()
}

func (json_object *data_t) Append_compile_time_var(new_var string) {
	json_object.Compile_Time_variables = append(json_object.Compile_Time_variables, new_var)
	json_object.Compile_Time_amount += 1
	json_object.Set_update_time()
}
func (json_object *data_t) Append_compile_time_value(new_value string) {
	json_object.Compile_Time_value = append(json_object.Compile_Time_value, new_value)
	json_object.Set_update_time()
}

func (json_object *data_t) Append_run_time_var(new_var string) {
	json_object.Run_Time_variables = append(json_object.Run_Time_variables, new_var)
	json_object.Run_Time_amount += 1
	json_object.Set_update_time()
}

func Create_object() data_t {
	var new_json data_t
	new_json.Set_creation_time() // Sets the time
	return new_json
}

func Convert_to_json(json_object data_t) []byte {
	serial_json, err := json.Marshal(json_object)
	if err != nil {
		notify.Error(err.Error(), "json.Convert_to_json()")
		var null []byte
		return null
	}
	return serial_json
}

func Convert_to_data_t(in_json []byte) data_t {
	var result data_t
	err := json.Unmarshal(in_json, &result)
	if err != nil {
		notify.Error(err.Error(), "json.Convert_to_data_t()")
		return data_t{}
	}
	return result
}

func Send(object data_t) string {
	return base64.StdEncoding.EncodeToString(Convert_to_json(object))
}

func Receive(object string) data_t {
	serialize_json, err := base64.StdEncoding.DecodeString(object)

	if err != nil {
		notify.Error(err.Error(), "json.Receive()")
	}

	return Convert_to_data_t(serialize_json)
}
