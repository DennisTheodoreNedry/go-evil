package tools

// Generates a random bool
func Generate_random_bool() bool {
	value := Generate_random_int_between(1, 2)

	return value == 1
}
