package structure

type Compile_var_t struct {
	Value string // The value that each variable contain
}

//
//
// Assigns the variable a value
//
//
func (object *Compile_var_t) Set_value(new_value string) {
	object.Value = new_value
}

//
//
// Grabs the variables value
//
//
func (object *Compile_var_t) Get_value() string {
	return object.Value
}
