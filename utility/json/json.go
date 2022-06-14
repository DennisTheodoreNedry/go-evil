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

	Target_OS   string `json:"Target_OS"`   // The OS you're targeting
	Target_ARCH string `json:"Target_ARCH"` // Target architecture
	Host_OS     string `json:"Host_OS"`     // The current running os

	Binary_name string `json:"BinaryName"` // The binary malware name
	Extension   string `json:"Extension"`  // Extension of the malware

	DebugMode bool `json:"DebugMode"` // Debug mode, doesn't delete the malware go file after compilation

	Verbose_LVL  string   `json:"VerLV"`       // Verbose level
	Call_history []string `json:"CallHistory"` // A string array containing all the functions we have passed through

	Malware_gut []string `json:"MalwareLine"` // The content of the malware

	File_Headers     []string `json:"UsedHeaders"`     // The encountered headers
	Imported_headers []string `json:"ImportedHeaders"` // The user specificed headers
	Disabled_domains []string `json:"DisabledHeaders"` // Headers that have been set to be disabled during this run

	Compile_Time_variables []string `json:"CompTimeVar"`    // The variables themself
	Compile_Time_value     []string `json:"CompTimeValue"`  // The values of the defined compile time variables
	Compile_Time_amount    int      `json:"CompTimeAmount"` // The amount of defined compile time variables

	Run_Time_variables []string `json:"RunTimeVar"`    // The variables themself
	Run_Time_amount    int      `json:"RunTimeAmount"` // The amount of defined run time variables

	Disable_region []string `json:"DisableRegion"` // The malware will not run on these regions

	Infection_count int `json:"Infection_Count"`

	Start_on_birth bool `json:"Start_on_birth"`
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

func (json_object *data_t) Append_imported_domain(new_header string) {

	for _, already_imported_header := range json_object.Imported_headers {
		if already_imported_header == new_header { // The header already exist, no need to add it
			return
		}
	}

	json_object.Imported_headers = append(json_object.Imported_headers, new_header)
	json_object.Set_update_time()
}

func (json_object *data_t) Get_imported_domain() []string {
	return json_object.Imported_headers
}

func (json_object *data_t) Is_imported(domain string) bool {
	toReturn := false

	for _, imported_domain := range json_object.Imported_headers {
		if imported_domain == domain { // Has this domain been imported?
			toReturn = true
		}
	}
	return toReturn
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

	for _, already_imported_header := range json_object.File_Headers {
		if already_imported_header == new_domain { // The header already exist, no need to add it
			return
		}
	}

	for _, disabled_domains := range json_object.Disabled_domains {
		if disabled_domains == new_domain { // The header is disabled, so ignore it.
			return
		}
	}

	json_object.File_Headers = append(json_object.File_Headers, new_domain)
	json_object.Set_update_time()
}

func (json_object *data_t) Get_File_domain() []string {
	return json_object.File_Headers
}

func (json_object *data_t) Append_malware_gut(new_line string) {
	json_object.Malware_gut = append(json_object.Malware_gut, new_line)
	json_object.Set_update_time()
}

func (json_object *data_t) Get_malware_gut() []string {
	return json_object.Malware_gut
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

func (json_object *data_t) Set_Extension(new_ext string) {
	json_object.Extension = new_ext
	json_object.Set_update_time()
}

func (json_object *data_t) Get_Extension() string {
	return json_object.Extension
}

func (json_object *data_t) Set_binary_name(new_name string) {
	json_object.Binary_name = new_name
	json_object.Set_update_time()
}
func (json_object *data_t) Get_binary_name() string {
	return json_object.Binary_name
}

func (json_object *data_t) Append_disabled_domain(new_header string) {
	for _, already_imported_header := range json_object.Disabled_domains {
		if already_imported_header == new_header { // The header already exist, no need to add it
			return
		}
	}

	json_object.Disabled_domains = append(json_object.Disabled_domains, new_header)
	json_object.Set_update_time()
}

func (json_object *data_t) Append_disabled_region(new_region string) {
	for _, line := range json_object.Disable_region {
		if line == new_region { // Region is already marked
			return
		}
	}

	json_object.Disable_region = append(json_object.Disabled_domains, new_region)
	json_object.Set_update_time()
}

func (json_object *data_t) Set_infect_count(new_count int) {
	json_object.Infection_count = new_count
	json_object.Set_update_time()
}

func (json_object *data_t) Get_infect_count() int {
	return json_object.Infection_count
}

func (json_object *data_t) Start_malware_on_birth(status bool) {
	json_object.Start_on_birth = status
	json_object.Set_update_time()
}

func (json_object *data_t) Get_status_start_malware_on_birth() bool {
	return json_object.Start_on_birth
}

// General functions
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
