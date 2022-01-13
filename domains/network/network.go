package network

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/s9rA16Bf4/go-evil/utility/contains"
	"github.com/s9rA16Bf4/go-evil/utility/variables/user"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type net_t struct {
	save_disk        bool                // If a Get request is performed, should the result be saved to disk?
	temp_file_prefix string              // Every created temp file will contain a prefix to mark them
	save_variable    bool                //  If a Get request is performed, should the result be saved to a runtime variable?
	data             map[string][]string // Header consists out of <key> <value>
	headers          []string
	latest_key       string
}

var c_net net_t

func POST(target_url string) {
	if len(c_net.latest_key) == 0 {
		notify.Error("No header was assigned", "network.POST()")
	}
	resp, err := http.PostForm(target_url, c_net.data) // Post
	if err != nil {
		notify.Error(err.Error(), "network.POST()")
	}
	user.Set_variable(resp.Status)
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
		notify.Error("Undefined header "+header, "network.POST_set_header()")
	}
	c_net.latest_key = header
	_, status := c_net.data[header]
	if !status {
		c_net.data[header] = make([]string, 0)
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
	}
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
		}
		body, _ := ioutil.ReadAll(resp.Body)
		dst.Write(body)
		user.Set_variable(dst.Name()) // Save the filename

	} else if c_net.save_variable {
		body, _ := ioutil.ReadAll(resp.Body)
		user.Set_variable(string(body))
	}
}

func GET_save_disk() {
	c_net.save_disk = true
}
func GET_save_variable() {
	c_net.save_variable = true
}
func GET_set_prefix(new_prefix string) {
	c_net.temp_file_prefix = new_prefix
}
