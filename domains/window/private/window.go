package domains

import (
	"github.com/s9rA16Bf4/go-evil/utility/converter"
	run_time "github.com/s9rA16Bf4/go-evil/utility/variables/runtime"
	"github.com/webview/webview"
)

type window struct {
	window_name string // Window name
	window_x    int    // Length on x axis
	window_y    int    // Length on y axis
}

var current_window window

// Functions that will not start the loop
func SetX(new_x string) {
	new_x = run_time.Check_if_variable(new_x)
	x := converter.String_to_int(new_x, "window.SetX()")
	current_window.window_x = x
}
func SetY(new_y string) {
	new_y = run_time.Check_if_variable(new_y)
	y := converter.String_to_int(new_y, "window.SetY()")
	current_window.window_y = y
}
func SetTitle(new_title string) {
	new_title = run_time.Check_if_variable(new_title)
	current_window.window_name = new_title
}

func preface() {
	if current_window.window_name == "" { // The user never told us what kind of name the window should utilize
		SetTitle("Untitled")
	}
	if current_window.window_y == 0 { // The user never specificed the length on the y axis
		SetY("400")
	}
	if current_window.window_x == 0 { // The user never specificed the length on the x axis
		SetX("200")
	}
}

// Functions that will start the loop
func GoToUrl(url string) {
	url = run_time.Check_if_variable(url)
	preface()
	win := webview.New(false)
	defer win.Destroy()
	win.SetTitle(current_window.window_name)
	win.SetSize(current_window.window_y, current_window.window_x, webview.HintNone)
	win.Navigate(url)
	win.Run()
}

func Display(msg string) {
	msg = run_time.Check_if_variable(msg)
	preface()
	win := webview.New(false)
	defer win.Destroy()
	win.SetTitle(current_window.window_name)
	win.SetSize(current_window.window_y, current_window.window_x, webview.HintNone)
	win.Navigate("data:text/html,<!doctype html><html><body><p>" + msg + "</p></body></html>")
	win.Run()
}
