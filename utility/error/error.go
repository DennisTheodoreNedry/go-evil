package error

import "github.com/s9rA16Bf4/notify_handler/go/notify"

func Subdomain_error(subdomain string, caller string) {
	notify.Error("Unknown subdomain '"+subdomain+"'", caller)
}
func Function_error(function string, caller string) {
	notify.Error("Unknown function '"+function+"'", caller)
}
