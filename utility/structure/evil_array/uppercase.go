package evilarray

import "strings"

// Makes the values in the array uppercase
func (object *Evil_array_t) Uppercase() {
	for i, value := range object.gut {
		object.Replace(i, strings.ToUpper(value))
	}
}
