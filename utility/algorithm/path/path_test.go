package path

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Generate_random_path(t *testing.T) {
	result := Generate_random_path(3)
	depth := strings.Count(result, "/")

	if depth-1 != 3 { // its whatever path the function generated + / at the beginning
		t.Log(fmt.Sprintf("Expected a depth of '3', recieved '%d'", depth))
		t.Fail()
	}
}

func Test_Generate_random_path_negative(t *testing.T) {
	result := Generate_random_path(-3)

	if result != "NULL" { // its whatever path the function generated + / at the beginning
		t.Log(fmt.Sprintf("Expected a result of 'NULL', recieved '%s'", result))
		t.Fail()
	}
}
