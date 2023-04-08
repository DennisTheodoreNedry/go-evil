package json

// Adds a global variable line to the final go code
func (object *Json_t) Add_go_global(new_global string) {

	for _, old := range object.GO_global {
		if old == new_global { // Check if the global already has been definied
			return
		}
	}

	object.GO_global = append(object.GO_global, new_global)
}
