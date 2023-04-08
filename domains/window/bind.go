package window

import "github.com/TeamPhoneix/go-evil/utility/structure"

// Binds a go function to a corresponding javascript function
func bind(values string, s_json string) string {
	data_structure := structure.Receive(s_json)
	arr := structure.Create_evil_object(values)

	js_call := arr.Get(0)
	evil_func := arr.Get(1)

	data_structure.Add_binding(js_call, evil_func)

	return structure.Send(data_structure)
}
