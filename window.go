package main

import (
	"github.com/webview/webview"
)

type window struct {
	window_name string // Window name
	window_x    int    // Length on x axis
	window_y    int    // Length on y axis
	window_text string // What is gonna be displayed inside the window
	url         string
}

var current_window window

// Functions that will not start the loop
func window_setText(new_text string) {
	current_window.window_text = new_text
}
func window_setX(new_x int) {
	current_window.window_x = new_x
}
func window_setY(new_y int) {
	current_window.window_y = new_y
}
func window_setTitle(new_title string) {
	current_window.window_name = new_title
}

func window_setDst(new_dst string) {
	current_window.url = new_dst
}

// Functions that will start the loop
func window_run() {
	if current_window.window_name == "" { // The user never told us what kind of name the window should utilize
		current_window.window_name = "Untitled"
	}
	if current_window.window_y == 0 { // The user never specificed the length on the y axis
		current_window.window_y = 400
	}
	if current_window.window_x == 0 { // The user never specificed the length on the x axis
		current_window.window_x = 200
	}

	win := webview.New(false)
	win.SetTitle(current_window.window_name)
	win.SetSize(current_window.window_y, current_window.window_x, webview.HintNone)
	win.Navigate(current_window.url)
	win.Run()
}
