package main

import (
	"fmt"
	"os"
)

var verbose_lvl = "0" // Default level, this will log nothing to the user

func notify_error(msg string, where string) { // Print msg and exit
	fmt.Println("#### notify_notify_error ####")
	fmt.Println("msg: " + msg)
	fmt.Println("where: " + where)
	os.Exit(1)
}

func notify_inform(msg string) {
	fmt.Println("[*] Inform: " + msg)
}

func notify_warning(msg string) {
	fmt.Println("[!] Warning: " + msg)
}

func notify_log(msg string, verbose_lvl string, suggested_verbose_lvl string) {
	if verbose_lvl != "0" && verbose_lvl >= suggested_verbose_lvl {
		fmt.Println("[%] " + msg)
	}
}
