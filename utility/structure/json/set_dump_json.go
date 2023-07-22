package json

// Print the json object after compilation
func (object *Json_t) Set_dump_json(value string) string {
	object.Dump_json = true
	return ""
}
