package tools

import (
	"math/rand"
	"time"
)

// Generates a random number between min and max
func Generate_random_int_between(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// Generates a random number between 1 and 128
func Generate_random_int() int {
	return Generate_random_int_between(1, 128)
}
