package domains

import (
	"github.com/s9rA16Bf4/go-evil/utility/converter"
	"github.com/webview/webview"
)

type window struct {
	window_name string // Window name
	window_x    int    // Length on x axis
	window_y    int    // Length on y axis
	window_text string // What is gonna be displayed inside the window
}

var current_window window

// Functions that will not start the loop
func Window_setText(new_text string) {
	current_window.window_text = new_text
}
func Window_setX(new_x string) {
	x := converter.String_to_int(new_x, "window.Window_setX()")
	current_window.window_x = x
}
func Window_setY(new_y string) {
	y := converter.String_to_int(new_y, "window.Window_setY()")
	current_window.window_y = y
}
func Window_setTitle(new_title string) {
	current_window.window_name = new_title
}

func window_preface() {
	if current_window.window_name == "" { // The user never told us what kind of name the window should utilize
		current_window.window_name = "Untitled"
	}
	if current_window.window_y == 0 { // The user never specificed the length on the y axis
		current_window.window_y = 400
	}
	if current_window.window_x == 0 { // The user never specificed the length on the x axis
		current_window.window_x = 200
	}
}

// Functions that will start the loop
func Window_goToUrl(url string) {
	window_preface()
	win := webview.New(false)
	defer win.Destroy()
	win.SetTitle(current_window.window_name)
	win.SetSize(current_window.window_y, current_window.window_x, webview.HintNone)
	win.Navigate(url)
	win.Run()
}

func Window_display(msg string) {
	window_preface()
	win := webview.New(false)
	defer win.Destroy()
	win.SetTitle(current_window.window_name)
	win.SetSize(current_window.window_y, current_window.window_x, webview.HintNone)
	win.Navigate("data:text/html,<!doctype html><html><body><p>" + msg + "</p></body></html>")
	win.Run()
}
