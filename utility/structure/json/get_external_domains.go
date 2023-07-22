package json

// Returns an array containing all externally defined domain paths
func (object *Json_t) Get_external_domains_path() []string {
	return object.External_domain_paths
}
