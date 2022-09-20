package webview

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
		<style>%s</style>
		<script>%s</script>
	</head>
	<body>
	%s
	</body>
	`, css_content, js_content, html_content)

	binding := ""
	for key := range data_object.Bind_gut {
		binding += fmt.Sprintf("w.Bind(%s, %s)\n", key, data_object.Bind_gut[key])
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", call),
		"w := webview.New(false)",
		"defer w.Destroy()",
		fmt.Sprintf("w.SetHtml(`%s`)", final_content),
		fmt.Sprintf("w.SetSize(%d, %d, webview.HintNone)", data_object.Width, data_object.Height),
		fmt.Sprintf("w.SetTitle(\"%s\")", data_object.Title),
		binding,
		"w.Run()",
		"}",
	})

	data_object.Add_go_import("github.com/webview/webview")

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

	data_object.Set_width(new_height)

	return structure.Send(data_object)
}

//
//
// Binds a go function to a corresponding javascript function
//
//
func bind(values string, s_json string) string {
	data_structure := structure.Receive(s_json)
	split_arr := tools.Extract_values_array(values)

	js_call := split_arr[0]
	evil_func := split_arr[1]

	data_structure.Add_binding(js_call, evil_func)

	return structure.Send(data_structure)
}

//
//
// Makes the webview enter a website of your choice
//
//
func navigate(website string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	call := "Navigate"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(website string){", call),
		"w := webview.New(false)",
		"defer w.Destroy()",
		fmt.Sprintf("w.SetSize(%d, %d, webview.HintNone)", data_object.Width, data_object.Height),
		fmt.Sprintf("w.SetTitle(\"%s\")", data_object.Title),
		"w.Navigate(website)",
		"w.Run()",
		"}",
	})

	data_object.Add_go_import("github.com/webview/webview")

	return []string{fmt.Sprintf("%s(%s)", call, website)}, structure.Send(data_object)
}
