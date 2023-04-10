package json

func (object *Json_t) Add_boot_function(new_function_call string) {
	object.Boot_functions = append(object.Boot_functions, new_function_call)
}
