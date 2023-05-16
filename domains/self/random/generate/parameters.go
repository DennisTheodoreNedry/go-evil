package generate

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/tools"
)

func Generate_parameters() ([]string, string) {
	amount_of_parameters := tools.Generate_random_int()
	sending_values := ""
	parameter_array := []string{}

	for y := 0; y < amount_of_parameters; y++ {
		parameter := fmt.Sprintf("param%d string", y)
		parameter_array = append(parameter_array, parameter)

		sending_values += fmt.Sprintf("\"%s\",", tools.Generate_random_string())
	}

	return parameter_array, sending_values
}
