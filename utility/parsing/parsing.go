package parsing

import (
	"github.com/TeamPhoneix/go-evil/utility/json"
)

//
// Parses the contents of the provided file
//
func Parse(s_json string) string {
	Check_for_erros(s_json)

	data_object := json.Receive(s_json)

	return json.Send(data_object)
}
