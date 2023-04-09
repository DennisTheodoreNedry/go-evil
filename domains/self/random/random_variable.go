package random

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Adds a random variable to the source code
func Add_variable(amount string, s_json string) string {
	data_object := structure.Receive(s_json)

	amount = tools.Erase_delimiter(amount, []string{"\""}, -1)

	i_value := tools.String_to_int(amount)
	if i_value == -1 {
		notify.Error(fmt.Sprintf("Unknown amount '%d'", i_value), "self.Add_random_variable()")
	}

	for i := 0; i < i_value; i++ {
		variable_name := tools.Generate_random_string()
		random_value := tools.Generate_random_string()

		data_object.Add_go_global(fmt.Sprintf("var %s string = \"%s\"", variable_name, random_value))
	}

	return structure.Send(data_object)
}
