package structure

import (
	"encoding/json"
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type json_t struct {
	//  File related stuff
	File_path string   `json:"file_path"`      // file path to the file we are reading
	File_gut  string   `json:"file_gut"`       // Contents of the file we read in
	Functions []Func_t `json:"file_functions"` // A structure containing all function strucs gathered

	// The malwares actual code will be found here
	GO_functions []string `json:"go_functions"` // Contains all the real go-code functions for the malware
	GO_imports   []string `json:"go_imports"`   // Contains all imports needed for the malware to work

	// Malware content
	Malware_gut      []string `json:"malware_gut"`     // The contents of the malware file
	Malware_Import   []string `json:"malware_imports"` // The libs the user wanted to include
	Malware_src_file string   `json:"malware_src"`     // The name of the malware src file
	Malware_path     string   `json:"malware_path"`    // The go file to compile

	// Malware configurations
	Target_os   string `json:"target_os"`   // The OS you are targeting
	Target_arch string `json:"target_arch"` // Target architecture
	Binary_name string `json:"binary_name"`
	Extension   string `json:"extension"`

	// Compiler configurations
	Debug_mode  bool   `json:"debug_mode"`
	Dump_json   bool   `json:"dump_json"`
	Obfuscate   bool   `json:"obfuscate"`
	Verbose_lvl string `json:"verbose_lvl"`

	// Text editor/webview
	Width    int               `json:"width"`         // The width of the text editor/webview
	Height   int               `json:"height"`        // The height of the text editor/webview
	Title    string            `json:"webview_title"` // Title of the webview window (not the text editor)
	Html_gut []string          `json:"html_gut"`      // The html code displayed in the webview (not the text editor)
	Js_gut   []string          `json:"js_gut"`        // The javascript code used in the webview (not the text editor)
	Css_gut  []string          `json:"css_gut"`       // The css code used in the webview (not the text editor)
	Bind_gut map[string]string `json:"bind_gut"`      // Contains all our bindings set by the user
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
// Adds all the contents of the read file
//
//
func (object *json_t) Add_file_gut(content string) {
	object.File_gut = content
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

	result := tools.Contains(ext, []string{"."}) // Checks if the extension contains a dot

	if status := result["."]; !status && ext != "" {
		ext = fmt.Sprintf(".%s", ext)
	}

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

//
//
// Sets how verbose the program should be
//
//
func (object *json_t) Set_verbose_lvl(value string) {
	object.Verbose_lvl = value
}

//
//
// Creates byte code from our json structure
//
//
func (object *json_t) Dump() []byte {
	serial_json, err := json.Marshal(object)

	if err != nil {
		notify.Error(err.Error(), "json.Convert_to_json()")
	}

	return serial_json
}

//
//
// Adds a function to the structure
//
//
func (object *json_t) Add_function(name string, f_type string, gut []string) {
	var new_func Func_t

	new_func.Set_name(name)

	if object.Obfuscate { // Doing it this way will also make sure that we populate the prevous names
		new_func.Set_name(tools.Generate_random_string())
	}

	new_func.Set_type(f_type)

	new_func.Add_lines(gut)

	object.Functions = append(object.Functions, new_func)
}

//
//
// Adds a domain to the imports
//
//
func (object *json_t) Add_domain(domain_name string) {
	object.Malware_Import = append(object.Malware_Import, domain_name)
}

//
//
// Adds a single line to the malware gut
//
//
func (object *json_t) Add_malware_line(line string) {
	object.Malware_gut = append(object.Malware_gut, line)
}

//
//
// Adds multiple lines to the malware gut
//
//
func (object *json_t) Add_malware_lines(lines []string) {
	object.Malware_gut = append(object.Malware_gut, lines...)
}

//
//
// Adds a go based function to the final go code
//
//
func (object *json_t) Add_go_function(lines []string) {
	function_call := lines[0]

	for _, calls := range object.GO_functions {
		if calls == function_call {
			return
		}
	}

	object.GO_functions = append(object.GO_functions, lines...)
}

//
//
// Adds a import line to the final go code
//
//
func (object *json_t) Add_go_import(new_import string) {

	new_import = fmt.Sprintf("\"%s\"", new_import)
	for _, old_import := range object.GO_imports {
		if old_import == new_import { // Check if the import already have been imported
			return
		}
	}

	object.GO_imports = append(object.GO_imports, new_import)
}

//
//
// Obfuscates the program
//
//
func (object *json_t) Enable_obfuscate() {
	object.Obfuscate = true
}

//
//
// Makes the source code readable
//
//
func (object *json_t) Disable_obfuscate() {
	object.Obfuscate = false
}

//
//
// Sets the width of the text editor
//
//
func (object *json_t) Set_width(value string) {
	if i_value := tools.String_to_int(value); i_value != -1 {
		object.Width = i_value
	} else {
		object.Width = 600
	}
}

//
//
// Sets the height of the text editor
//
//
func (object *json_t) Set_height(value string) {
	if i_value := tools.String_to_int(value); i_value != -1 {
		object.Height = i_value
	} else {
		object.Height = 800
	}
}

//
//
// Sets the html code being displayed
//
//
func (object *json_t) Set_html(content string) {
	content = tools.Erase_delimiter(content, `"`)
	object.Html_gut = append(object.Html_gut, content)
}

//
//
// Sets the js code being used in the webview
//
//
func (object *json_t) Set_js(content string) {
	content = tools.Erase_delimiter(content, `"`)
	object.Js_gut = append(object.Js_gut, content)
}

//
//
// Sets the css code being used in the webview
//
//
func (object *json_t) Set_css(content string) {
	content = tools.Erase_delimiter(content, `"`)
	object.Css_gut = append(object.Css_gut, content)
}

//
//
// Sets the height of the text editor
//
//
func (object *json_t) Set_title(value string) {
	object.Title = value
}

//
//
// Adds a binding to the webview
// it's accessible by running `window.<js_call>()` in your html code
//
//
func (object *json_t) Add_binding(js_call string, evil_call string) {
	evil_call = tools.Erase_delimiter(evil_call, `"`)

	if object.Obfuscate { // Assuming that we need to convert the name to the obfuscated version
		for _, d_func := range object.Functions {
			if d_func.Check_previous_names(evil_call) {
				evil_call = d_func.Name
				break
			}
		}
	}

	object.Bind_gut[js_call] = evil_call
}
