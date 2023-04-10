package json

func (object *Json_t) Add_loop_function(new_function_call string) {
	object.Loop_functions = append(object.Loop_functions, new_function_call)
}
