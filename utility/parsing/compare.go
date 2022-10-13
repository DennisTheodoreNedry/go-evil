package parsing

import (
	"fmt"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

//
//
// Construcs the code needed to be able to compare
//
//
func construct_compare(condition string, s_json string) ([]string, string) {
	call := []string{"compare"}

	data_object := structure.Receive(s_json)
	arr := structure.Create_evil_object(condition)

	final_body := []string{fmt.Sprintf(
		"func %s(values []string){", call[0]),
		"for _, value := range values{",
		"value = spine.variable.get(value)",
		"spine.variable.foreach = value"}

	final_body = append(final_body, "}}")

	data_object.Add_go_function(final_body)

	call[0] = fmt.Sprintf("%s(%s)", call[0], arr.To_string("array"))

	return call, structure.Send(data_object)
}
