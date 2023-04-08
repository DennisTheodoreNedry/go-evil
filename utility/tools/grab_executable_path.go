package tools

import "os"

// Returns the path to the executable file
// including the executable file name
func Grab_executable_path() string {
	path := os.Args[0]
	return path
}
