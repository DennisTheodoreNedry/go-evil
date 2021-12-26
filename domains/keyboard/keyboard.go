package keyboard

import (
	"os/exec"
	"regexp"
	"runtime"
)

func Keyboard_lock() {
	// Which os are we running on?
	// We need to lock the keyboard and then the mouse
	switch runtime.GOOS {
	case "linux":
		possible_arguments := []string{"xinput", "xtrlock"}

		for _, arg := range possible_arguments {
			if arg == "xinput" {
				out, err := exec.Command("xinput", "list").Output()
				if err == nil { // xinput existed!
					regex := regexp.MustCompile("([0-9]+)") // This grabs all numbers found, which are devices
					ids := regex.FindAllStringSubmatch(string(out[:]), -1)
					for _, id := range ids {
						exec.Command("xinput", "float", id[1]).Run() // Just loop through and turn everything off
					}
				}
			} else if arg == "xtrlock" {
				exec.Command("xtrlock").Run()
			}
		}

	case "windows":

	case "darwin":
	}
}

func Keyboard_unlock() {
	switch runtime.GOOS {
	case "linux":
		possible_arguments := []string{"xinput", "xtrlock"}
		for _, arg := range possible_arguments {

			if arg == "xinput" {
				out, err := exec.Command("xinput", "list").Output()
				if err == nil { // xinput existed!
					regex := regexp.MustCompile("\\[slave +keyboard \\(([0-9]+)\\)\\]")
					master_id := regex.FindAllStringSubmatch(string(out[:]), -1)[0][1] // Grab the first case of 'slave keyboard'

					regex = regexp.MustCompile("([0-9]+)")
					ids := regex.FindAllStringSubmatch(string(out[:]), -1)
					for _, id := range ids {
						exec.Command("xinput", "reattach", id[1], master_id).Run() // Just loop through and turn everything on, don't really care if they get connected to the same device again
					}
				}
			} else if arg == "xtrlock" {
				exec.Command("killall", "xtrlock").Run() // If this doesn't work well then welp
			}
		}

	case "windows":

	case "darwin":
	}
}
