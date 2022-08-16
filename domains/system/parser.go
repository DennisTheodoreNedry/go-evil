package system

import (
	"github.com/TeamPhoneix/go-evil/utility/json"
)

//
//
// The main parser for the system domain
//
//
func Parser(s_json string) string {
	data_object := json.Receive(s_json)

	return json.Send(data_object)
}
