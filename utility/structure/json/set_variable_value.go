package json

// Sets the value of a compiletime variable
func (object *Json_t) Set_variable_value(value string) {
	object.Comp_var[object.Comp_id].Set_value(value)
	object.Comp_id++

	if object.Comp_id >= object.Var_max { // Reset
		object.Comp_id = 0
	}
}
