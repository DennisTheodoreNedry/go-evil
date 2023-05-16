package until

import (
	"fmt"
	"regexp"
	"time"

	"github.com/s9rA16Bf4/go-evil/utility/structure/functions"
	"github.com/s9rA16Bf4/go-evil/utility/structure/json"
	"github.com/s9rA16Bf4/go-evil/utility/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

const (
	GRAB_FULL_DATE = "([0-9]{0,4})/([0-9]{0,2})/([0-9]{0,2})-([0-9]{0,2}):([0-9]{0,2})" // YYYY/MM/DD-hh:mm
	GRAB_HOUR_MIN  = "([0-9]{0,2}):([0-9]{0,2})"                                        // hh:mm
)

// Makes the malware wait until the yyyy-mm-dd-hh-mm has been reached
func Until(value string, data_object *json.Json_t) []string {
	function_call := "Until"

	regex := regexp.MustCompile(GRAB_FULL_DATE)
	result_full := regex.FindAllStringSubmatch(value, -1)

	regex2 := regexp.MustCompile(GRAB_HOUR_MIN)
	result_hour_min := regex2.FindAllStringSubmatch(value, -1)

	if len(result_full) == 0 && len(result_hour_min) == 0 {
		notify.Error("Failed to find a valid time format!", "time.Until")
	}

	year, month, day := time.Now().Date()

	parameters := []string{
		tools.Int_to_string(year),
		tools.Int_to_string(int(month)),
		tools.Int_to_string(day),
		// hour
		// minute
	}

	if len(result_full) > 0 {
		parameters = []string{result_full[0][1], result_full[0][2], result_full[0][3], result_full[0][4], result_full[0][5]}
	} else {
		parameters = append(parameters, result_hour_min[0][1], result_hour_min[0][2]) // Grab the hour and minute
	}

	data_object.Add_go_function(functions.Go_func_t{Name: function_call, Func_type: "", Part_of_struct: "", Return_type: "",
		Parameters: []string{"repr_1 []int", "repr_2 []int", "repr_3 []int", "repr_4 []int", "repr_5 []int"},

		Gut: []string{
			"year := tools.String_to_int(spine.variable.get(spine.alpha.construct_string(repr_1)))",
			"month := tools.String_to_int(spine.variable.get(spine.alpha.construct_string(repr_2)))",
			"day := tools.String_to_int(spine.variable.get(spine.alpha.construct_string(repr_3)))",
			"hour := tools.String_to_int(spine.variable.get(spine.alpha.construct_string(repr_4)))",
			"minute := tools.String_to_int(spine.variable.get(spine.alpha.construct_string(repr_5)))",
			"c_now := time.Now()",
			"for year != c_now.Year() || month != int(c_now.Month()) || day != c_now.Day() || hour != c_now.Hour() || minute != c_now.Minute() {",
			"time.Sleep(5 * (10 ^ 9))",
			"c_now = time.Now()",
			"}",
		}})

	data_object.Add_go_import("time")

	constructed_parameters := []string{
		data_object.Generate_int_array_parameter(parameters[0]),
		data_object.Generate_int_array_parameter(parameters[1]),
		data_object.Generate_int_array_parameter(parameters[2]),
		data_object.Generate_int_array_parameter(parameters[3]),
		data_object.Generate_int_array_parameter(parameters[4]),
	}

	return []string{fmt.Sprintf("%s(%s, %s, %s, %s, %s)", function_call, constructed_parameters[0], constructed_parameters[1], constructed_parameters[2],
		constructed_parameters[3], constructed_parameters[4])}
}
