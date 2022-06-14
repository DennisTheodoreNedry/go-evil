package parser

import (
	"fmt"
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/json"
	"github.com/s9rA16Bf4/go-evil/utility/version"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	EXTRACT_MAIN_FUNC_HEADER    = "(main:{)"                          // We use this to identify if there are multiple main functions in the same file
	EXTRACT_FUNCTION_CALL       = "([@#a-z]+).*\\((.*)\\);"           // Grabs function and a potential value
	EXTRACT_FUNCTION_CALL_WRONG = "([@#a-z]+).*\\((\"(.*)\")?\\)[^;]" // And this is utilized to find rows that don't end in ;
	EXTRACT_COMPILER_VERSION    = "\\[ ?version ([0-9]+\\.[0-9]+)\\]" // Extracts the major version
)

type data_t struct {
	main_method        int      // Adds one for each main method found, should always be 1.
	compiler_flag      bool     // Is the compiler version listed
	forgotten_semi     []string // Contains all the rows were a ; has been forgotten
	forgotten_semi_row []int    // Contains the row numbers
}

func (data *data_t) Flag_main_method() {
	data.main_method += 1
}
func (data *data_t) Get_main_method_amount() int {
	return data.main_method
}

func (data *data_t) Flag_compiler() {
	data.compiler_flag = true
}
func (data *data_t) Get_compiler_flag() bool {
	return data.compiler_flag
}

func (data *data_t) Forgot_semi(line string, row int) {
	data.forgotten_semi = append(data.forgotten_semi, line)
	c_data_t.forgotten_semi_row = append(c_data_t.forgotten_semi_row, row)
}
func (data *data_t) Forgot_semi_amount() int {
	return len(data.forgotten_semi)
}

var c_data_t data_t

func compiler_version(line string) {
	if !c_data_t.Get_compiler_flag() {
		regex := regexp.MustCompile(EXTRACT_COMPILER_VERSION) // Extracts the high and medium version
		compiler_version := regex.FindAllStringSubmatch(line, -1)
		if len(compiler_version) > 0 {
			listed_version := compiler_version[0][1]
			if version.Get_Compiler_version() < listed_version {
				notify.Error("Unknown compiler version "+listed_version, "parser.Parser()")
			} else if version.Get_Compiler_version() > listed_version {
				notify.Warning("You're running a script for an older version of the compiler.\nThis means that there might be functions/syntaxes that have changed")
			}
			c_data_t.Flag_compiler()
		}
	}
}

func find_main(line string) {
	regex := regexp.MustCompile(EXTRACT_MAIN_FUNC_HEADER)
	main_function := regex.FindAllStringSubmatch(line, -1)
	if len(main_function) == 1 {
		c_data_t.Flag_main_method()
	}
}

func find_semicolon(line string, row int) {
	regex := regexp.MustCompile(EXTRACT_FUNCTION_CALL_WRONG)
	match := regex.FindAllStringSubmatch(line, -1)

	if len(match) > 0 {
		c_data_t.Forgot_semi(line, row)
	}
}

func Look_for_errors(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("parser.Look_for_errors()")

	for i, line := range data_structure.File_gut {
		find_main(line)
		compiler_version(line)
		find_semicolon(line, i)
	}

	if c_data_t.Get_main_method_amount() == 0 {
		notify.Error("Failed to find a main function in the provided file '"+data_structure.File+"'",
			"parser.Look_for_errors()")
	} else if c_data_t.Get_main_method_amount() > 1 {
		notify.Error("Found multiple main definitions in the provided file '"+data_structure.File+"'",
			"parser.Look_for_errors()")
	}

	if !c_data_t.Get_compiler_flag() {
		notify.Error("No major version was specificed", "parser.Look_for_errors()")
	}

	if c_data_t.Forgot_semi_amount() > 0 {
		for i, line := range c_data_t.forgotten_semi {
			notify.Error(fmt.Sprintf("There is a missing semi-colon on row %d, '%s' ", c_data_t.forgotten_semi_row[i]+1, line),
				"parser.Look_for_errors()")
		}
	}

	return json.Send(data_structure)
}
