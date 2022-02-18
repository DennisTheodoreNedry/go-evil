package compiler_time

import (
	"regexp"
	"strings"

	"github.com/s9rA16Bf4/go-evil/utility/converter"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type var_t struct {
	variable        [5]string // $1 - $5
	index           int       // Where we currently are in the array above
	latest_variable int       // Which was the latest variable that we modified
}

var curr_var var_t

const (
	EXTRACT_VARIABLE = "\\$[0-9]+"
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
		found_value := "NULL"
		variable := result[0][0]
		var_int := converter.String_to_int(variable[1:], "compilter_time.Get_variable()")
		if var_int > 0 || var_int < 6 {
			if curr_var.variable[var_int-1] != "" {
				found_value = curr_var.variable[var_int-1]
			}
		} else {
			notify.Error("Illegal index "+variable, "compilter_time.Get_variable()")
			return index

		}
		index = strings.Replace(index, variable, found_value, 1)
		notify.Log("Found variable "+variable+" which contained the value "+found_value, notify.Verbose_lvl, "2")
	}
	return index
}

func Get_latest_variable() int {
	return curr_var.latest_variable
}
