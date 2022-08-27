package tools

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/thanhpk/randstr"
)

//
//
// Wrapper for the strings.HasPrefix function, but takes in a array contaning strings to look for
// and returns a map in the format of { "<key>":"<true/false>" }
//
//
func Starts_with(target string, selection []string) map[string]bool {
	to_return := make(map[string]bool)

	for _, value := range selection {
		to_return[value] = strings.HasPrefix(target, value)
	}

	return to_return
}

//
//
// Wrapper for the strings.HasSuffix function, but takes in a array contaning strings to look for
// and returns a map in the format of { "<key>":"<true/false>" }
//
//
func Ends_with(target string, selection []string) map[string]bool {
	to_return := make(map[string]bool)

	for _, value := range selection {
		to_return[value] = strings.HasSuffix(target, value)
	}
	return to_return
}

//
//
// Wrapper for the strings.Contains function, but takes in a array contaning strings to look for
// and returns a map in the format of { "<key>":"<true/false>" }
//
//
func Contains(target string, selection []string) map[string]bool {
	to_return := make(map[string]bool)

	for _, value := range selection {
		to_return[value] = strings.Contains(target, value)
	}
	return to_return
}

//
// Tries to convert the provided string to an int, and returns either
// the converted value or -1 if it failed
//
func String_to_int(value string) int {
	i_value, err := strconv.Atoi(value) // Tries to convert

	if err != nil {
		return -1
	}

	return i_value
}

//
//
// Converts the provided int to a string
//
//
func Int_to_string(value int) string {
	return fmt.Sprint(value)
}

//
//
// Takes the string to work in as a pointer, and makes it go from abc to cba
//
//
func Reverse_string(target *string) {
	output := make([]rune, utf8.RuneCountInString(*target))
	roof := len(output)

	for _, character := range *target {
		roof--
		output[roof] = character
	}
	*target = string(output[0:])
}

//
//
// Splits a string into a string array and returns it
//
//
func Split_string(target string) []string {
	to_return := []string{}
	chars := []rune(target)
	for i := 0; i < len(chars); i++ {
		to_return = append(to_return, string(chars[i]))
	}
	return to_return
}

//
//
// Generates a random string based on the length of the input
//
//
func Generate_random_string(size int) string {
	var toReturn string
	for {
		toReturn = randstr.String(size)
		if _, err := strconv.Atoi(string(toReturn[0])); err != nil { // It can't start with a number
			break
		}
	}
	return toReturn
}
