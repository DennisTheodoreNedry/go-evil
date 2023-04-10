package json

func (object *Json_t) Add_end_function(new_function_call string) {
	object.End_functions = append(object.End_functions, new_function_call)
}
