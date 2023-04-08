package json

// Adds a go based function to the final go code
func (object *Json_t) Add_go_function(lines []string) {
	function_call := lines[0]

	for _, calls := range object.GO_functions {
		if calls == function_call {
			return
		}
	}

	object.GO_functions = append(object.GO_functions, lines...)
}
