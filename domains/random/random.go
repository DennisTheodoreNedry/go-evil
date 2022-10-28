package random

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

//
//
// Generates a random int value in a range
// The input is an evil array that should only contain two values, min and max
// The generated value is placed in a compile-time variable
//
//
func generate_int(value string, s_json string) string {
	data_object := structure.Receive(s_json)
	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Expected an array with two positions, got %d", arr.Length()), "random.generate_int()")
	}

	min := arr.Get(0)
	i_min := tools.String_to_int(min)

	if i_min == -1 {
		notify.Error(fmt.Sprintf("Failed to convert min value %s to an integer!", min), "random.generate_int()")
	}

	max := arr.Get(1)
	i_max := tools.String_to_int(max)

	if i_max == -1 {
		notify.Error(fmt.Sprintf("Failed to convert max value %s to an integer!", max), "random.generate_int()")
	}

	generated_value := tools.Generate_random_int_between(i_min, i_max)

	data_object.Set_variable_value(tools.Int_to_string(generated_value))

	return structure.Send(data_object)
}

//
//
// Generates a random string where the provded value is the length of the string
// The generated value is placed in a compile-time variable
//
//
func generate_string(value string, s_json string) string {
	data_object := structure.Receive(s_json)

	roof := tools.Erase_delimiter(value, []string{"\""}, -1)
	length := tools.String_to_int(roof)

	if length == -1 {
		notify.Error(fmt.Sprintf("Failed to convert value %s to an integer!", roof), "random.generate_string()")
	}

	generated_value := tools.Generate_random_n_string(length)

	data_object.Set_variable_value(generated_value)

	return structure.Send(data_object)
}
