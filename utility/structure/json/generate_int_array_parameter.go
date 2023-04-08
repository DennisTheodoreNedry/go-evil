package json

import "fmt"

// Generates a string which in turn represents an int array based on the input message
func (object *Json_t) Generate_int_array_parameter(message string) string {
	to_return := "[]int{"
	for _, repr := range object.Generate_int_array(message) {
		to_return += fmt.Sprintf("%d,", repr)
	}
	to_return += "}"

	return to_return
}
