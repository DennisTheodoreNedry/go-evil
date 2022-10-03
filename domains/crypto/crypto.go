package crypto

import "github.com/TeamPhoneix/go-evil/utility/structure"

//
//
// Encrypts the provided object
// The input must follow this format '${"<encryption system>", "<key>", "<file/ext/dir>", "<path/to/object>"}$'
//
//
func encrypt(value string, s_json string) string {
	data_object := structure.Receive(s_json)

	return structure.Send(data_object)
}

func set_method(value string, s_json string) string {
	data_object := structure.Receive(s_json)

	return structure.Send(data_object)
}

func set_key(value string, s_json string) string {
	data_object := structure.Receive(s_json)

	return structure.Send(data_object)
}

func set_target(value string, s_json string) string {
	data_object := structure.Receive(s_json)

	return structure.Send(data_object)
}

func decrypt(value string, s_json string) string {
	data_object := structure.Receive(s_json)

	return structure.Send(data_object)
}
