package notify

import (
	"fmt"
	"os"
)

var Verbose_lvl = "0" // Default level, this will log nothing to the user

func Notify_error(msg string, where string) { // Print msg and exit
	fmt.Println("#### Error ####")
	fmt.Println("msg: " + msg)
	fmt.Println("where: " + where)
	os.Exit(1)
}

func Notify_inform(msg string) {
	fmt.Println("[*] Inform: " + msg)
}

func Notify_warning(msg string) {
	fmt.Println("[!] Warning: " + msg)
}

func Notify_log(msg string, verbose_lvl string, suggested_verbose_lvl string) {
	if verbose_lvl != "0" && verbose_lvl >= suggested_verbose_lvl {
		fmt.Println("[%] " + msg)
	}
}
