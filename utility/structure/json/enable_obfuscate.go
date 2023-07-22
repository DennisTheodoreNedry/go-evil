package json

// Obfuscates the program
func (object *Json_t) Enable_obfuscate(value string) string {
	object.Obfuscate = true
	return ""
}
