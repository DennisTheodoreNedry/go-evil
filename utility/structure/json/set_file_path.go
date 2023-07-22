package json

// Updates the internal file path
func (object *Json_t) Set_file_path(new_path string) string {
	object.File_path = new_path
	return ""
}
