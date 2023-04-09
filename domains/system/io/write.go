package io

import (
	"fmt"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

// Writes a provided content to a provided file
func Write(s_json string, value string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Write"

	arr := structure.Create_evil_object(value)
	path := arr.Get(0)
	data := strings.Join(arr.Get_between(1, arr.Length()), " ")

	if data_object.Check_global_var(data) { // Checks if what we got is a global variable
		data = tools.Erase_delimiter(data, []string{"\""}, -1)
	} else {
		data = fmt.Sprintf("\"%s\"", data)
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1  []int, repr_2 []int){", function_call),
		"path := spine.alpha.construct_string(repr_1)",
		"path = spine.variable.get(path)",

		"content := spine.alpha.construct_string(repr_2)",
		"content = spine.variable.get(content)",

		"file, err := os.Create(path)",
		"if err != nil{",
		"spine.log(err.Error())",
		"return",
		"}",

		"defer file.Close()",
		"result := tools.Starts_with(content, []string{\"[HEX];\"})",
		"if ok := result[\"[HEX];\"]; !ok {",
		"file.WriteString(content)",
		"}else{",
		"split := strings.Split(content, \",\")",
		"for _, data := range split {",
		"data, _ := hex.DecodeString(data)",
		"file.Write(data)",
		"}}}",
	})

	data_object.Add_go_import("encoding/hex")
	data_object.Add_go_import("os")
	data_object.Add_go_import("strings")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/tools")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	parameter_path := data_object.Generate_int_array_parameter(path)
	parameter_data := data_object.Generate_int_array_parameter(data)

	return []string{fmt.Sprintf("%s(%s, %s)", function_call, parameter_path, parameter_data)}, structure.Send(data_object)
}