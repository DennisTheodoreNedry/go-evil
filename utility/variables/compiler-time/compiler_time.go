package compiler_time

import (
	"github.com/s9rA16Bf4/go-evil/utility/converter"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type var_t struct {
	variable        [5]string // $1 - $5
	index           int       // Where we currently are in the array above
	latest_variable int       // Which was the latest variable that we modified
}

var curr_var var_t

func Set_variable(value string) {
	if curr_var.index >= 5 {
		curr_var.index = 0
	}
	curr_var.latest_variable = curr_var.index // Remember
	curr_var.variable[curr_var.index] = value // Add the value
	curr_var.index++                          // Index to the next address
}

func Get_variable(index string) string {
	index = index[1:]                                                    // Gets whatever is left after $
	id := converter.String_to_int(index, "compiler_time.Get_variable()") // Convert

	if id > 5 || id < 1 {
		notify.Error("Out-of-bonds variable "+index, "compiler_time.Get_variable()")
		return "NULL"
	}

	if curr_var.variable[id-1] == "" { // If it's empty
		return "NULL"
	}
	return curr_var.variable[id-1]
}

func Get_latest_variable() int {
	return curr_var.latest_variable
}
