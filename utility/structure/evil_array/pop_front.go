package evilarray

// Removes the value in the front and returns it
func (object *Evil_array_t) Pop_front() string {
	to_return := object.gut[0]

	// Move everything one step back
	for i := 1; i < len(object.gut); i++ {
		object.gut[i-1] = object.gut[i]
	}

	object.gut[len(object.gut)-1] = "" // Remove the old element
	object.length--

	return to_return
}
