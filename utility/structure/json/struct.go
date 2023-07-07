package json

import (
	compilevar "github.com/s9rA16Bf4/go-evil/utility/structure/compile_var"
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/structs"
	notify "github.com/s9rA16Bf4/notify_handler"
)

type Json_t struct {
	// Build directory
	Build_directory string `json:"build_directory"` // The directory where everything will be placed

	//  File related stuff
	File_path string             `json:"file_path"`      // file path to the file we are reading
	File_gut  string             `json:"file_gut"`       // Contents of the file we read in
	Functions []functions.Func_t `json:"file_functions"` // A structure containing all functions def in the files

	// The malwares actual code will be found here
	GO_functions []functions.Go_func_t `json:"go_functions"` // Contains all the real go-code functions for the malware
	GO_imports   []string              `json:"go_imports"`   // Contains all imports needed for the malware to work
	GO_const     []string              `json:"go_const"`     // Contains all consts needed for the malware to work
	GO_global    []string              `json:"go_global"`    // Contains all globals needed for the malware to work
	GO_struct    []structs.Go_struct_t `json:"go_struct"`    // Contains all structs

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

	// Text editor/window
	Width    int               `json:"width"`        // The width of the text editor/window
	Height   int               `json:"height"`       // The height of the text editor/window
	Title    string            `json:"window_title"` // Title of the window window (not the text editor)
	Html_gut []string          `json:"html_gut"`     // The html code displayed in the window (not the text editor)
	Js_gut   []string          `json:"js_gut"`       // The javascript code used in the window (not the text editor)
	Css_gut  []string          `json:"css_gut"`      // The css code used in the window (not the text editor)
	Bind_gut map[string]string `json:"bind_gut"`     // Contains all our bindings set by the user

	// Variables
	Var_max  int                        `json:"variable_max"`      // The max amount of allowed variables
	Comp_var []compilevar.Compile_var_t `json:"compile_variables"` // All the compile time variables
	Comp_id  int                        `json:"compile_ID"`        // The current index for the compile variable

	// Debugger behavior
	Debugger_behavior string `json:"debugger_behavior"` // How should the malware behave after detecting a debugger being used?

	// Custom alphabet
	Alphabet []string `json:"alphabet"` // Internal alphabet utilized

	// Function call order
	Boot_functions []string `json:"boot_function"`
	Loop_functions []string `json:"loop_function"`
	End_functions  []string `json:"end_function"`

	// External domain (lib) folders
	External_domain_paths []string `json:"external_domain_paths"`

	// Log object
	Log_object notify.Verbose_t `json:"log_object`
}
