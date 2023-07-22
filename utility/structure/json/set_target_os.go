package json

// Sets the target os for the compiler
func (object *Json_t) Set_target_os(os string) string {
	if os == "windows" && object.Extension == "" {
		object.Set_extension(".exe")
	}
	object.Target_os = os

	return ""
}
