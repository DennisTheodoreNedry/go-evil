package evilarray

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Grabs data at the provided index
func (object *Evil_array_t) Get(index int) string {
	if index < 0 || index > object.Length() {
		notify.Error(fmt.Sprintf("Index %d was out-of-bound", index), "evil_array.Replace()")
	}
	return object.gut[index]
}

// Grabs all data between the provided indexes
func (object *Evil_array_t) Get_between(start int, end int) []string {
	to_return := []string{}
	if start < 0 || start > object.Length() {
		notify.Error(fmt.Sprintf("Index %d was out-of-bound", start), "evil_array.Replace()")
	}

	if end < 0 || end > object.Length() {
		notify.Error(fmt.Sprintf("Index %d was out-of-bound", end), "evil_array.Replace()")
	}

	for i := start; i < end; i++ {
		to_return = append(to_return, object.gut[i])
	}

	return to_return
}
