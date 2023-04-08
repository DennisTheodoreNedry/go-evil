package evilarray

// Removes the value at the index and returns it
func (object *Evil_array_t) Pop_index(index int) string {
	to_return := object.Get(index)
	object.gut[index] = "" // Remove the old element

	// Move everything one step back
	for i := index; i < len(object.gut); i++ {
		object.gut[i-1] = object.gut[i]
	}

	object.length--

	return to_return
}
