package tools

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/thanhpk/randstr"
)

// Generates a random string based on the length of the input
func Generate_random_n_string(size int) string {
	var toReturn string
	for {
		toReturn = randstr.String(size)
		if _, err := strconv.Atoi(string(toReturn[0])); err != nil { // It can't start with a number
			break
		}
	}
	return toReturn
}

// Generates a random string between the lengthes of 1 and 128
func Generate_random_string() string {
	rand.Seed(time.Now().UnixNano())
	max := 128
	min := 1
	return Generate_random_n_string(rand.Intn(max-min) + min)
}
