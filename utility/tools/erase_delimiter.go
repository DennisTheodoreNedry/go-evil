package tools

import "strings"

// Erases all occurences of the delimiter in the string
func Erase_delimiter(line string, delimiters []string, count int) string {

	for _, delimiter := range delimiters {
		line = strings.Replace(line, delimiter, "", count)
	}

	return line
}
