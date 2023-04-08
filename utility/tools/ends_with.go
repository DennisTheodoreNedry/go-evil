package tools

import "strings"

// Wrapper for the strings.HasSuffix function, but takes in a array contaning strings to look for
// and returns a map in the format of { "<key>":"<true/false>" }
func Ends_with(target string, selection []string) map[string]bool {
	to_return := make(map[string]bool)

	for _, value := range selection {
		to_return[value] = strings.HasSuffix(target, value)
	}
	return to_return
}
