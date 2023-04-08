package tools

import "unicode/utf8"

// Takes the string to work in as a pointer, and makes it go from abc to cba
func Reverse_string(target *string) {
	output := make([]rune, utf8.RuneCountInString(*target))
	roof := len(output)

	for _, character := range *target {
		roof--
		output[roof] = character
	}
	*target = string(output[0:])
}
