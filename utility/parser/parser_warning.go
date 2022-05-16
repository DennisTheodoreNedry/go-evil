package parser

import "github.com/s9rA16Bf4/notify_handler/go/notify"

func development_warning() {
	notify.Warning("Development flag was set to true, essentially the project is not perfect and usage may veary")
}
