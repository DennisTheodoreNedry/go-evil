package runtime

import (
	"fmt"
	"os/user"
	"regexp"
	"strings"

	"github.com/s9rA16Bf4/go-evil/utility/converter"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type var_t struct {
	variable        [5]string // €1 - €5
	index           int       // Where we currently are in the array above
	latest_variable int       // Which was the latest variable that we modified

	static_index           int          // Where we currently are in the static array
	static_variables       [10][]string // Static variables that range from $10 - $20
	latest_static_variable int          // Last modified static variable
}

var curr_var var_t

const (
	EXTRACT_VARIABLE = "€[0-9]+"
)

func Set_variable(value string) {
	if curr_var.index >= 5 {
		curr_var.index = 0
	}
	curr_var.latest_variable = curr_var.index // Remember
	curr_var.variable[curr_var.index] = value // Add the value
	curr_var.index++                          // Index to the next address
}

func Get_variable(index string) string {
	regex := regexp.MustCompile(EXTRACT_VARIABLE)
	result := regex.FindAllStringSubmatch(index, -1)
	if len(result) >= 1 {
		found_value := ""
		variable := result[0][0]
		var_int := converter.String_to_int(variable[3:], "runtime.Get_variable()")
		if var_int == 666 {
			user, err := user.Current()
			if err != nil {
				notify.Error(err.Error(), "runtime.Get_variable()")
				return index
			}
			found_value = user.Name
		} else if var_int > 0 && var_int < 6 {
			found_value = curr_var.variable[var_int-1]
		} else {
			if var_int > 9 && var_int < 21 { // Range of static variables
				for _, line := range curr_var.static_variables[var_int-10] {
					if line != "" { // It contains something
						found_value = strings.Join(curr_var.static_variables[var_int-10], "\n")
						break
					}
				}
			} else {
				notify.Error(fmt.Sprintf("Illegal index %s", variable), "runtime.Get_variable()")
				return index
			}
		}
		index = strings.Replace(index, variable, found_value, 1)
	}
	return index
}

func Get_latest_variable() int {
	return curr_var.latest_variable
}

func Get_latest_value() string {
	return curr_var.variable[Get_latest_variable()]
}

func Check_if_variable(input string) string {
	var variable string
	var start_of_variable = false
	for _, c := range input {
		if string(c) == "€" {
			start_of_variable = true
		} else if string(c) == " " && start_of_variable {
			break
		}
		if start_of_variable {
			variable += string(c)
		}
	}
	if Get_variable(variable) != "" {
		input = strings.Replace(input, variable, Get_variable(variable), 1)
	}

	return input
}

func Set_Static_Variable(new_value []string) {
	if curr_var.static_index >= 10 {
		curr_var.static_index = 0
	}
	curr_var.latest_static_variable = curr_var.static_index
	curr_var.static_variables[curr_var.static_index] = new_value
	curr_var.static_index++ // Index to the next address
}

func Get_latest_static_variable() int {
	return curr_var.latest_static_variable
}
