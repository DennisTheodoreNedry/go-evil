package json

// Adds a domain to the imports
func (object *Json_t) Add_domain(domain_name string) {
	object.Malware_Import = append(object.Malware_Import, domain_name)
}
