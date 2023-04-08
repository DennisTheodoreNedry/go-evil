package json

import "fmt"

// Returns a string consisting of the internal alphabet
func (object *Json_t) Get_alphabet() string {
	to_return := "[]string{"

	for _, repr := range object.Alphabet {
		to_return += fmt.Sprintf("\"%s\",", repr)
	}

	to_return += "}"

	return to_return
}
