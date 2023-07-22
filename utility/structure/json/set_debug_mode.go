package json

// Sets the debug mode that the compiler will obey
func (object *Json_t) Set_debug_mode(mode string) string {
	if mode == "false" {
		object.Debug_mode = false
	} else {
		object.Debug_mode = true
	}

	return ""
}
