package parsing

import (
	"fmt"
	"regexp"

	"github.com/TeamPhoneix/go-evil/utility/json"
)

//
//
// Checks for invalidity in the provided file
//
//
func Check_for_erros(s_json string) {
	check_func(s_json)

}

//
//
// Checks if the required functions exist
//
//
func check_func(s_json string) {
	data_object := json.Receive(s_json)

	regex := regexp.MustCompile(FUNC)

	//required_functions := map[string]bool{"init": false, "infect": false, "loop": false}

	for _, line := range data_object.File_gut {
		result := regex.FindAllStringSubmatch(line, -1)

		fmt.Println(line, result)

	}
}
