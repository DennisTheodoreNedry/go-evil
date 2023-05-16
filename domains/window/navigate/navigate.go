package navigate

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
)

// Makes the window enter a website of your choice
func Navigate(website string, data_object *json.Json_t) []string {
	function_call := "Navigate"

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int"},

		Gut: []string{
			"website := spine.variable.get(spine.alpha.construct_string(repr_1))",
			fmt.Sprintf("win, err := lorca.New(website, \"\",%d, %d)", data_object.Width, data_object.Height),
			"if err != nil{",
			"spine.log(err.Error())",
			"return",
			"}",
			"defer win.Close()",
			"<-win.Done()",
		}})

	data_object.Add_go_import("github.com/zserge/lorca")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	parameter_1 := data_object.Generate_int_array_parameter(website)

	return []string{fmt.Sprintf("%s(%s)", function_call, parameter_1)}
}
