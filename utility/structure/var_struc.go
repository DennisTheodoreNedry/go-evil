package structure

type Var_t struct {
	Value string // The value that each variable contain
}

//
//
// Assigns the internal `value` variabel to a value
//
//
func (object *Var_t) Set_value(new_value string) {
	object.Value = new_value
}

//
//
// Grabs the value for this variable
//
//
func (object *Var_t) Get_value() string {
	return object.Value
}
