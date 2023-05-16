package network

import (
	"fmt"
	"strings"

	"github.com/s9rA16Bf4/go-evil/utility/structure"
	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Pings a target, takes in an evil array with the following contents
// ${'target', 'count', 'udp/tcp'}$
// target - ip or addr
// count - How many times to ping, 0 for indefinitely
// udp/tcp - Which protocol do you want to use?

func Ping(value string, data_object *json.Json_t) []string {
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

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int", "count int", "repr_2 []int"},
		Gut: []string{
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
		}})

	data_object.Add_go_import("fmt")
	data_object.Add_go_import("net")
	data_object.Add_go_import("time")
	data_object.Add_go_import("github.com/tatsushid/go-fastping")
	data_object.Add_go_import("github.com/s9rA16Bf4/notify_handler/go/notify")

	int_target := data_object.Generate_int_array_parameter(target)
	int_protocol := data_object.Generate_int_array_parameter(protocol)

	return []string{fmt.Sprintf("%s(%s, %d, %s)", function_call, int_target, count, int_protocol)}
}
