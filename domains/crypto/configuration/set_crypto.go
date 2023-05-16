package configuration

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Sets the crypto system to use for encrypting and decrypting
func Set_crypto(value string, data_object *json.Json_t) []string {
	available_systems := []string{"aes", "rsa"}
	function_call := "set_crypto"
	value = tools.Erase_delimiter(value, []string{"\""}, -1)

	def_crypto := false // Is the crypto that we are gonna use definied?

	for _, def_c := range available_systems {
		if def_c == value {
			def_crypto = true
			break
		}
	}

	if !def_crypto { // Failed to find the crypto
		notify.Error(fmt.Sprintf("Unknown crypto system '%s', available are %s", value, available_systems), "crypto.set_method()")
	}

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int"},
		Gut: []string{
			"target := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"spine.crypt.set_crypto(target)",
		}})

	// Construct our int array
	parameter := data_object.Generate_int_array_parameter(value)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter)}
}
