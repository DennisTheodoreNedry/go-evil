package tools

import "path/filepath"

// Returns the executable name
func Grab_executable_name() string {
	path := Grab_executable_path()
	exe_name := filepath.Base(path)

	return exe_name
}
