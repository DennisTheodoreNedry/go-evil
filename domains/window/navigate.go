package window

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Makes the window enter a website of your choice
func navigate(website string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Navigate"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int){", function_call),
		"website := spine.variable.get(spine.alpha.construct_string(repr_1))",
		fmt.Sprintf("win, err := lorca.New(website, \"\",%d, %d)", data_object.Width, data_object.Height),
		"if err != nil{",
		"spine.log(err.Error())",
		"return",
		"}",
		"defer win.Close()",
		"<-win.Done()",
		"}",
	})

	data_object.Add_go_import("github.com/zserge/lorca")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	parameter_1 := data_object.Generate_int_array_parameter(website)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}, structure.Send(data_object)
}
