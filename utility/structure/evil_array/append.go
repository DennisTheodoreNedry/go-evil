package evilarray

// Appends data to the array
func (object *Evil_array_t) Append(new_content string) {
	object.gut = append(object.gut, new_content)
	object.length++
}
