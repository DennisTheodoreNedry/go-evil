package arguments

import (
	"fmt"
	"os"

	"github.com/s9rA16Bf4/go-evil/utility/notify"
)

type argument struct {
	longName          string   // Long command for this argument, i.e. '--help'
	shortName         string   // Short command for this argument, i.e. '-h' instead of '--help'
	desc              string   // Description of the command
	argument_required bool     // Does this command require a value to function?
	argument_value    string   // This will contain the passed value if something was passed
	set               bool     // Has this argument been passed?
	options           []string // Available options for this command
}

var defined_arguments []argument // This array will contain all the defined arguments

func Argument_add(longName string, shortName string, argument_required bool, desc string, options []string) {
	// Adds an argument
	defined_arguments = append(defined_arguments, argument{longName, shortName, desc, argument_required, "NULL", false, options})
}
func Argument_parse() {
	// Parses all entered arguments, check if they contain the correct value and does all magic
	for i := 1; i <= len(os.Args[1:]); i++ {
		found := false
		for j, def_arg := range defined_arguments {
			if os.Args[i] == def_arg.longName || os.Args[i] == def_arg.shortName {
				if def_arg.argument_required && i+1 > len(os.Args[1:]) {
					notify.Notify_error("notify_error: The argument "+os.Args[i]+" needs an argument to work!", "arguments.argument_parse()")
				} else if def_arg.argument_required {
					possible_options := false
					for _, opt := range def_arg.options {
						if opt == os.Args[i+1] || opt == "NULL" {
							defined_arguments[j].argument_value = os.Args[i+1] // Save the value
							i += 1
							possible_options = true
							break
						}
					}
					if !possible_options {
						notify_error_msg := "notify_error: Unknown options " + os.Args[i+1] + " to argument " + os.Args[i] + ", possible options are ["
						for i, option := range def_arg.options {
							notify_error_msg += option
							if i+1 < len(def_arg.options) {
								notify_error_msg += ","
							}
						}
						notify_error_msg += "]"

						notify.Notify_error(notify_error_msg, "arguments.argument_parse()")
					}
				}

				defined_arguments[j].set = true
				found = true
				break
			}
		}
		if !found {
			notify.Notify_error("notify_error: The argument "+os.Args[i]+" was not defined!", "arguments.argument_parse()")
		}
	}
}

func Argument_check(arg_name string) bool {
	// Checks if a argument has been called
	var toReturn bool = false
	for _, arg := range defined_arguments {
		if arg.longName == arg_name || arg.shortName == arg_name {
			toReturn = arg.set
			break
		}
	}
	return toReturn
}

func Argument_get(arg_name string) string {
	// Gets the provided value for a set argument, and will return NULL otherwise
	var toReturn string = "NULL"
	for _, arg := range defined_arguments {
		if (arg.longName == arg_name || arg.shortName == arg_name) && arg.set {
			toReturn = arg.argument_value
			break
		}
	}
	return toReturn
}

func Argument_help() {
	// Prints all defined arguments
	fmt.Println("#### Definied Arguments ####")
	for id := 0; id < len(defined_arguments); id++ {
		fmt.Print(defined_arguments[id].longName + ", " + defined_arguments[id].shortName)

		if defined_arguments[id].argument_required {
			fmt.Print(" <value> ")
		}
		fmt.Println(" | " + defined_arguments[id].desc)
	}
}
