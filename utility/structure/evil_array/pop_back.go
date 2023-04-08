package evilarray

// Removes the value in the back and returns it
func (object *Evil_array_t) Pop_back() string {
	to_return := object.gut[len(object.gut)-1]

	object.gut[len(object.gut)-1] = "" // Remove the old element

	object.length--

	return to_return
}
