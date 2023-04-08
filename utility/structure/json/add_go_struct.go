package json

// Adds a structure to the final code
func (object *Json_t) Add_go_struct(new_struct []string) {
	struct_header := new_struct[0]

	for _, old_header := range object.GO_struct {
		if old_header == struct_header { // Check if the struct already has been definied
			return
		}
	}

	object.GO_struct = append(object.GO_struct, new_struct...)
}
