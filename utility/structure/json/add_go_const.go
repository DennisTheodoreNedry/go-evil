package json

// Adds a const line to the final go code
func (object *Json_t) Add_go_const(new_const string) {
	for _, old := range object.GO_const {
		if old == new_const { // Check if the const already has been definied
			return
		}
	}

	object.GO_const = append(object.GO_const, new_const)
}
