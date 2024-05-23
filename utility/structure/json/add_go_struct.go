package json

import "github.com/DennisTheodoreNedry/go-evil/utility/structure/structs"

// Adds a structure to the final code
func (object *Json_t) Add_go_struct(new_struct structs.Go_struct_t) {

	for _, old_struct := range object.GO_struct {
		if old_struct.Name == new_struct.Name { // Check if the struct already has been definied
			return
		}
	}

	object.GO_struct = append(object.GO_struct, new_struct)
}
