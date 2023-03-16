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
	target := arr.Get(0)

	count := tools.String_to_int(arr.Get(1))
	if count == -1 {
		notify.Error(fmt.Sprintf("Failed to convert '%s' to an integer", arr.Get(1)), "network.ping()")
	}

	protocol := strings.ToLower(arr.Get(2))

	if arr.Length() != 3 {
		notify.Error(fmt.Sprintf("Expected three values, but recieved %d", arr.Length()), "network.ping()")
	}

	data_object.Add_go_function([]string{
		fmt.Sprintf("func %s(repr_1 []int, count int, repr_2 []int){", function_call),
		"target := spine.variable.get(spine.alpha.construct_string(repr_1))",
		"protocol := spine.variable.get(spine.alpha.construct_string(repr_2))",

		"pinger := fastping.NewPinger()",
		"ra, err := net.ResolveIPAddr(\"ip4:icmp\", target)",
		"if err != nil {",
		"notify.Log(err.Error(), spine.logging, \"3\")",
		"}",
		"pinger.AddIPAddr(ra)",
		"switch (protocol){",
		"case \"tcp\":",
		"pinger.Network(\"tcp\")",
		"case \"udp\":",
		"pinger.Network(\"udp\")",
		"default:",
		"notify.Log(fmt.Sprintf(\"Unknown protocol %s\", protocol), spine.logging, \"3\")",
		"}",
		"pinger.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {",
		"spine.variable.set(fmt.Sprintf(\"IP Addr: %s receive, RTT: %v\", addr.String(), rtt))",
		"}",
		"if count == 0{",
		"for {",
		"err = pinger.Run()",
		"if err != nil {",
		"notify.Log(err.Error(), spine.logging, \"3\")",
		"}",
		"}",
		"} else{",
		"for i := 0; i < count; i++ {",
		"err = pinger.Run()",
		"if err != nil {",
		"notify.Log(err.Error(), spine.logging, \"3\")",
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
