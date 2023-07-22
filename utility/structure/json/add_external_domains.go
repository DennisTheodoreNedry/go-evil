package json

// Adds an external domain which the program can use
func (object *Json_t) Add_external_domains_path(folder_path string) string {
	object.External_domain_paths = append(object.External_domain_paths, folder_path)

	return ""
}
