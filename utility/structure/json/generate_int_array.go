package json

import gotools "github.com/DennisTheodoreNedry/Go-tools"

// Generates an int array representing the provided string
func (object *Json_t) Generate_int_array(message string) []int {
	to_return := []int{}

	message = gotools.EraseDelimiter(message, []string{"\""}, -1)

	for _, c_msg := range message {
		for id, c_alpha := range object.Alphabet {
			if string(c_msg) == string(c_alpha) {
				to_return = append(to_return, id)
			}
		}
	}

	return to_return
}
