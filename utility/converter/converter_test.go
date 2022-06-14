package converter

import (
	"fmt"
	"testing"
)

func Test_String_to_int(t *testing.T) {
	result := String_to_int("3", "converter.Test_String_to_int()")
	if result != 3 {
		t.Log(fmt.Sprintf("Result should have been 3, got %d", result))
		t.Fail()
	}

	result = String_to_int("-3", "converter.Test_String_to_int()")
	if result != -3 {
		t.Log(fmt.Sprintf("Result should have been -3, got %d", result))
		t.Fail()
	}
}
