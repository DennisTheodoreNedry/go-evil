package strip

func Strip(s_json string) string {
	s_json = remove_comments(s_json)
	s_json = remove_configuration(s_json)
	s_json = remove_imports(s_json)

	return s_json
}
