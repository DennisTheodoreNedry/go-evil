package functions

import "strings"

// Adds a line of content to the functions gut
func (object *Func_t) Add_line(line string) {
	object.Gut = append(object.Gut, strings.Split(line, "\n")...) // Split it up in newlines so that we don't lose content
}

// Adds multiple lines of content to the functions gut
func (object *Func_t) Add_lines(lines []string) {
	for _, line := range lines {
		object.Add_line(line)
	}
}
