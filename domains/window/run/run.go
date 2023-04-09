package run

import (
	"fmt"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

// Drops all the needed code from the json strucutre into one function
func Run(s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "Run()"

	html_content := strings.Join(data_object.Html_gut, "\n")
	js_content := strings.Join(data_object.Js_gut, "\n")
	css_content := strings.Join(data_object.Css_gut, "\n")

	final_content := fmt.Sprintf(`
	<head>
		<title>%s</title>
		<style>%s</style>
		<script>%s</script>
	</head>
	<body>
	%s
	</body>
	`, data_object.Title, css_content, js_content, html_content)

	binding := ""
	for key := range data_object.Bind_gut {
		binding += fmt.Sprintf("w.Bind(%s, %s)\n", key, data_object.Bind_gut[key])
	}

	body := []string{
		fmt.Sprintf("func %s(){", function_call),
		fmt.Sprintf("win, err := lorca.New(fmt.Sprintf(\"data:text/html,%%s\", url.PathEscape(`%s`)), \"\", %d, %d)", final_content, data_object.Width, data_object.Height),
		"if err != nil{",
		"spine.log(err.Error())",
		"return",
		"}",
	}

	for key := range data_object.Bind_gut {
		body = append(body, fmt.Sprintf("win.Bind(%s, %s)\n", key, data_object.Bind_gut[key]))
	}

	body = append(body, "defer win.Close()", "<-win.Done()", "}")

	data_object.Add_go_function(body)

	data_object.Add_go_import("github.com/zserge/lorca")
	data_object.Add_go_import("net/url")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")
	data_object.Add_go_import("fmt")

	return []string{function_call}, structure.Send(data_object)
}
