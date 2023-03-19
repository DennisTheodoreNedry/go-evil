package window

import (
	"fmt"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
)

//
//
// Drops all the needed code from the json strucutre into one function
//
//
func run(s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	call := "Run()"

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
		fmt.Sprintf("func %s(){", call),
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

	return []string{call}, structure.Send(data_object)
}

//
//
// Sets the html content displayed
//
//
func set_html(html_content string, s_json string) string {
	data_object := structure.Receive(s_json)

	data_object.Set_html(html_content)

	return structure.Send(data_object)
}

//
//
// Sets the js that wil be used
//
//
func set_js(js_content string, s_json string) string {
	data_object := structure.Receive(s_json)

	data_object.Set_js(js_content)

	return structure.Send(data_object)
}

//
//
// Sets the css that will be used
//
//
func set_css(css_content string, s_json string) string {
	data_object := structure.Receive(s_json)

	data_object.Set_css(css_content)

	return structure.Send(data_object)
}

//
//
// Sets the title of the window that appears
//
//
func set_title(new_title string, s_json string) string {
	data_object := structure.Receive(s_json)
	new_title = tools.Erase_delimiter(new_title, []string{"\""}, -1)

	data_object.Set_title(new_title)

	return structure.Send(data_object)
}

//
//
// Sets the width of the window
//
//
func set_width(new_width string, s_json string) string {
	data_object := structure.Receive(s_json)
	new_width = tools.Erase_delimiter(new_width, []string{"\""}, -1)

	data_object.Set_width(new_width)

	return structure.Send(data_object)
}

//
//
// Sets the height of the window
//
//
func set_height(new_height string, s_json string) string {
	data_object := structure.Receive(s_json)
	new_height = tools.Erase_delimiter(new_height, []string{"\""}, -1)

	data_object.Set_height(new_height)

	return structure.Send(data_object)
}

//
//
// Binds a go function to a corresponding javascript function
//
//
func bind(values string, s_json string) string {
	data_structure := structure.Receive(s_json)
	arr := structure.Create_evil_object(values)

	js_call := arr.Get(0)
	evil_func := arr.Get(1)

	data_structure.Add_binding(js_call, evil_func)

	return structure.Send(data_structure)
}

//
//
// Makes the window enter a website of your choice
//
//
func navigate(website string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	call := "Navigate"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(website string){", call),
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

	return []string{fmt.Sprintf("%s(%s)", call, website)}, structure.Send(data_object)
}
