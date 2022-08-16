package json

type json_t struct {
	File_path string   `json:"file_path"` // file path to the file we are reading
	File_gut  []string `json:"file_gut"`  // Contents of the file we read in

	Malware_gut      []string `json:"malware_gut"`  // The contents of the malware file
	Malware_src_file string   `json:"malware_src"`  // The name of the malware src file
	Malware_path     string   `json:"malware_path"` // The go file to compile

	Target_os   string `json:"target_os"`   // The OS you are targeting
	Target_arch string `json:"target_arch"` // Target architecture

	Binary_name string `json:"binary_name"`
	Extension   string `json:"extension"`

	Debug_mode bool `json:"debug_mode"` // Debug mode
	Dump_json  bool `json:"Dump_json"`

	Verbose_lvl string `json:"verbose_lvl"`
}

//
//
// Updates the internal file path
//
//
func (object *json_t) Set_file_path(new_path string) {
	object.File_path = new_path
}

//
//
// Appends a single line to the internal file gut
//
//
func (object *json_t) Add_file_gut(new_line string) {
	object.File_gut = append(object.File_gut, new_line)
}

//
//
// Appends a multiple lines to the internal file gut
//
//
func (object *json_t) Add_file_guts(new_lines []string) {
	object.File_gut = append(object.File_gut, new_lines...)
}

//
//
// Sets the target os for the compiler
//
//
func (object *json_t) Set_target_os(os string) {
	object.Target_os = os
}

//
//
// Sets the target arch for the compiler
//
//
func (object *json_t) Set_target_arch(arch string) {
	object.Target_arch = arch
}

//
//
// Sets the binaries name
//
//
func (object *json_t) Set_binary_name(name string) {
	object.Binary_name = name
}

//
//
// Sets the the extension
//
//
func (object *json_t) Set_extension(ext string) {
	object.Extension = ext
}

//
//
// Sets the debug mode that the compiler will obey
//
//
func (object *json_t) Set_debug_mode(mode string) {
	if mode == "false" {
		object.Debug_mode = false
	} else {
		object.Debug_mode = true
	}
}

//
//
// Print the json object after compilation
//
//
func (object *json_t) Set_dump_json() {
	object.Dump_json = true
}

func (object *json_t) Set_verbose_lvl(value string) {
	object.Verbose_lvl = value
}
