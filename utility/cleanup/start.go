package cleanup

// Wrapper function calling
// * remove_src_file
// * dump_json
func Start(s_json string) {
	remove_src_file(s_json)
	dump_json(s_json)
}
