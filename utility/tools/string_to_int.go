package tools

import "strconv"

// Tries to convert the provided string to an int, and returns either
// the converted value or -1 if it failed
func String_to_int(value string) int {
	i_value, err := strconv.Atoi(value) // Tries to convert

	if err != nil {
		return -1
	}

	return i_value
}
