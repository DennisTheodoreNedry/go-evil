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
func run(s_json string) (string, string) {
	data_object := structure.Receive(s_json)
	call := "run()"
	obj := "w"

	if data_object.Obfuscate {
		call = fmt.Sprintf("%s()", tools.Generate_random_string())
		obj = tools.Generate_random_string()
	}

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
		binding += fmt.Sprintf("%s.Bind(%s, %s)\n", obj, key, data_object.Bind_gut[key])
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", call),
		fmt.Sprintf("%s := webview.New(false)", obj),
		fmt.Sprintf("defer %s.Destroy()", obj),
		fmt.Sprintf("%s.SetHtml(`%s`)", obj, final_content),
		fmt.Sprintf("%s.SetSize(%d, %d, webview.HintNone)", obj, data_object.Width, data_object.Height),
		fmt.Sprintf("%s.SetTitle(\"%s\")", obj, data_object.Title),
		binding,
		fmt.Sprintf("%s.Run()", obj),
		"}",
	})

	data_object.Add_go_import("github.com/webview/webview")

	return call, structure.Send(data_object)
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
func navigate(website string, s_json string) (string, string) {
	data_object := structure.Receive(s_json)
	call := "navigate"
	obj := "w"

	if data_object.Obfuscate {
		call = fmt.Sprintf("%s()", tools.Generate_random_string())
		obj = tools.Generate_random_string()
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", call),
		fmt.Sprintf("%s := webview.New(false)", obj),
		fmt.Sprintf("defer %s.Destroy()", obj),
		fmt.Sprintf("%s.SetSize(%d, %d, webview.HintNone)", obj, data_object.Width, data_object.Height),
		fmt.Sprintf("%s.SetTitle(\"%s\")", obj, data_object.Title),
		fmt.Sprintf("%s.Navigate(%s)", obj, website),
		fmt.Sprintf("%s.Run()", obj),
		"}",
	})

	data_object.Add_go_import("github.com/webview/webview")

	return call, structure.Send(data_object)
}
