package evilarray

import "strings"

// Makes the values in the array lowercase
func (object *Evil_array_t) Lowercase() {
	for i, value := range object.gut {
		object.Replace(i, strings.ToLower(value))
	}
}
