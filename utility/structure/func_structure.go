package structure

import "strings"

type Func_t struct {
	Name      string   // Function name
	Func_type string   // Which sort of type this function is
	Gut       []string // The contents of the function
}

//
//
// Gives the function a name
//
//
func (object *Func_t) Set_name(name string) {
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

	for _, delimiter := range []string{"\n", "\t"} { // Cleans up the data
		line = strings.Split(line, delimiter)[0]
	}

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
