package contains

import (
	"regexp"

	"github.com/s9rA16Bf4/go-evil/utility/reverse"
)

const (
	EXTRACT_FUNCTION_VALUE = ".+\\(\"(.*)\"\\);" // Grabs the value being passed to the function
)

func StartsWith(string_to_look_in string, what_to_look_for []string) bool {
	toReturn := false

	if string_to_look_in == "" {
		return false
	}

	for _, tar := range what_to_look_for {
		input_string := []rune(string_to_look_in)
		if tar != "" { // If not empty
			for i, c := range tar {
				if c == input_string[i] {
					toReturn = true
				} else {
					toReturn = false
					break
				}
			}
		}

		if toReturn {
			break
		}
	}
	return toReturn
}

func EndsWith(string_to_look_in string, what_to_look_for []string) bool {
	reverse.ReverseString(&string_to_look_in) // This should reverse our string
	var value []string
	for _, target := range what_to_look_for {
		reverse.ReverseString(&target) // This helps us alot
		value = append(value, target)
	}

	return StartsWith(string_to_look_in, value)
}

func Contains(string_to_look_in string, string_to_look_for string) bool {
	toReturn := false

	for i, _ := range string_to_look_in {
		if i+1 < len(string_to_look_in) {
			d := i
			for _, c := range string_to_look_for {
				if d+1 > len(string_to_look_in) {
					break
				} else {
					if c == rune(string_to_look_in[d]) {
						toReturn = true
					} else {
						toReturn = false
						break
					}
					d++
				}
			}
		}
		if toReturn { // We have already found what we were after
			break
		}
	}

	return toReturn
}

func Passed_value(line string) string {
	regex := regexp.MustCompile(EXTRACT_FUNCTION_VALUE)
	result := regex.FindAllStringSubmatch(line, -1)
	var value string
	if len(result) > 0 {
		value = result[0][1]
	} else {
		value = "NULL"
	}
	return value
}
