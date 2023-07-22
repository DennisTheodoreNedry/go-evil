package evilarray

// Returns the inner gut which is where all the arrays contents is located
func (object *Evil_array_t) Dump() []string {
	return object.gut
}
