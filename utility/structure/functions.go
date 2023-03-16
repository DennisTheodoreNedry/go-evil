package structure

import (
	"fmt"
	"strings"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type Func_t struct {
	Name        string   // Obfuscated function name or the real function name
	Func_type   string   // Which sort of type this function is
	Return_type string   // What kind of value can we expect to be returned
	Gut         []string // The contents of the function
}

//
//
// Gives the function a name
//
//
func (object *Func_t) Set_name(name string) {

	for _, illegal_name := range []string{"main"} {
		if illegal_name == name {
			notify.Error(fmt.Sprintf("Illegal name '%s' was found!", name), "func_structure.Set_name()")
		}
	}

	object.Name = name
}

//
//
// Sets which type of function this is
//
//
func (object *Func_t) Set_type(f_type string) {
	object.Func_type = f_type
}

//
//
// Adds a line of content to the functions gut
//
//
func (object *Func_t) Add_line(line string) {

	for _, line := range strings.Split(line, "\n") { // Split it up in newlines so that we don't lose content

		found := true

		for _, g_line := range object.Gut { // Check for duplicates

			if g_line == line {
				found = false
				break
			}

		}

		if found {
			object.Gut = append(object.Gut, line)
		}
	}

}

//
//
// Adds multiple lines of content to the functions gut
//
//
func (object *Func_t) Add_lines(lines []string) {
	for _, line := range lines {
		object.Add_line(line)
	}
}

//
//
// Returns the functions name
//
//
func (object *Func_t) Get_name() string {
	return object.Name
}

//
//
// Returns the functions type
//
//
func (object *Func_t) Get_type() string {
	return object.Func_type
}

//
//
// Returns the functions content
//
//
func (object *Func_t) Get_gut() []string {
	return object.Gut
}

//
//
// Sets the return value
//
//
func (object *Func_t) Set_return_type(value string) {
	object.Return_type = value
}
