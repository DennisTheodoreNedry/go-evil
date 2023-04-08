package json

import "fmt"

// Adds a import line to the final go code
func (object *Json_t) Add_go_import(new_import string) {

	new_import = fmt.Sprintf("\"%s\"", new_import)
	for _, old_import := range object.GO_imports {
		if old_import == new_import { // Check if the import already have been imported
			return
		}
	}

	object.GO_imports = append(object.GO_imports, new_import)
}
