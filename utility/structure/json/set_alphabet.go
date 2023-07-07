package json

import "strings"

// Sets the internal alphabet utilized by the malware
func (object *Json_t) Set_alphabet(alphabet string) string {
	object.Alphabet = nil // Reset

	for _, char := range strings.Split(alphabet, ",") {
		object.Alphabet = append(object.Alphabet, string(char))
	}
	return ""
}
