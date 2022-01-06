package domains

import (
	"github.com/s9rA16Bf4/go-evil/utility/converter"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
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
	x := converter.String_to_int(new_x, "window.SetX()")
	notify.Log("Setting the length of the x axis to "+new_x, notify.Verbose_lvl, "3")
	current_window.window_x = x
}
func SetY(new_y string) {
	y := converter.String_to_int(new_y, "window.SetY()")
	notify.Log("Setting the length of the y axis to "+new_y, notify.Verbose_lvl, "3")
	current_window.window_y = y
}
func SetTitle(new_title string) {
	notify.Log("Setting the window title to "+new_title, notify.Verbose_lvl, "3")
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
	notify.Log("Will set the target url to "+url, notify.Verbose_lvl, "3")
	preface()
	win := webview.New(false)
	defer win.Destroy()
	win.SetTitle(current_window.window_name)
	win.SetSize(current_window.window_y, current_window.window_x, webview.HintNone)
	win.Navigate(url)
	win.Run()
}

func Display(msg string) {
	notify.Log("Will display the message "+msg, notify.Verbose_lvl, "3")
	preface()
	win := webview.New(false)
	defer win.Destroy()
	win.SetTitle(current_window.window_name)
	win.SetSize(current_window.window_y, current_window.window_x, webview.HintNone)
	win.Navigate("data:text/html,<!doctype html><html><body><p>" + msg + "</p></body></html>")
	win.Run()
}
