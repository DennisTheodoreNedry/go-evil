package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/converter"
	user "github.com/s9rA16Bf4/go-evil/utility/variables/runtime"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type net_t struct {
	save_disk        bool                // If a Get request is performed, should the result be saved to disk?
	temp_file_prefix string              // Every created temp file will contain a prefix to mark them
	save_variable    bool                //  If a Get request is performed, should the result be saved to a runtime variable?
	data             map[string][]string // Header consists out of <key> <value>
	headers          []string            // All defined headers
	latest_key       string              // The latest defined header
	ping_roof        string              // How many icmp packages should be sent
	ping_success     int                 // 0 - false, 1 - true
}

var c_net net_t

func POST(target_url string) {
	if len(c_net.latest_key) == 0 {
		notify.Error("No header was assigned", "network.POST()")
	} else {
		resp, err := http.PostForm(target_url, c_net.data) // Post
		if err != nil {
			notify.Error(err.Error(), "network.POST()")
		} else {
			user.Set_variable(resp.Status)
		}
	}

}
func POST_add_header(new_header string) {
	if c_net.data == nil {
		c_net.data = make(map[string][]string)
	}
	c_net.headers = append(c_net.headers, new_header)
}
func POST_set_header(header string) {
	found := false
	for _, defined_header := range c_net.headers {
		if defined_header == header {
			found = true
			break
		}
	}
	if !found {
		notify.Error(fmt.Sprintf("Undefined header %s", header), "network.POST_set_header()")
	} else {
		c_net.latest_key = header
		_, status := c_net.data[header]
		if !status {
			c_net.data[header] = make([]string, 0)
		}
	}

}
func POST_bind_value_to_latest_header(value string) {
	c_net.data[c_net.latest_key] = append(c_net.data[c_net.latest_key], value)
}

func GET(target_url string) {
	target_url = user.Check_if_variable(target_url)

	if !contains.StartsWith(target_url, []string{"http://", "https://"}) {
		target_url = "https://" + target_url
	}
	resp, err := http.Get(target_url)

	if err != nil {
		notify.Error(err.Error(), "network.Get()")
	} else {
		if c_net.save_disk { // Save the result to the disk
			if c_net.temp_file_prefix == "" {
				GET_set_prefix("eat_my_ass-")
			}
			if !contains.EndsWith(c_net.temp_file_prefix, []string{"-"}) {
				c_net.temp_file_prefix += "-"
			}

			dst, err := ioutil.TempFile(os.TempDir(), c_net.temp_file_prefix)
			if err != nil {
				notify.Error(err.Error(), "network.Get()")
			} else {
				body, _ := ioutil.ReadAll(resp.Body)
				dst.Write(body)
				user.Set_variable(dst.Name()) // Save the filename
			}

		} else if c_net.save_variable {
			body, _ := ioutil.ReadAll(resp.Body)
			user.Set_variable(string(body))
		}
	}

}

func GET_save_disk() { // Saves the result to disk
	c_net.save_disk = true
}
func GET_save_variable() { // Saves the result to a runtime variable
	c_net.save_variable = true
}
func GET_set_prefix(new_prefix string) { // Tells the system what should the output file be named before a random string is applied
	c_net.temp_file_prefix = new_prefix
}

func Ping(target string) {
	if c_net.ping_roof == "" {
		Ping_set_roof("5")
	}
	out, err := exec.Command("ping", target, fmt.Sprintf("-c %s", c_net.ping_roof)).Output()
	if err != nil {
		c_net.ping_success = 0
		notify.Error(err.Error(), "network.Ping()")
	} else {
		c_net.ping_success = 1
		user.Set_variable(string(out)) // Save the result
	}

}

func Ping_set_roof(new_roof string) {
	if converter.String_to_int(new_roof, "network.Ping_set_roof()") == -1 { // Checks if it works
		return
	}
	c_net.ping_roof = new_roof
}
