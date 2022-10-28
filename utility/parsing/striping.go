package parsing

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/TeamPhoneix/go-evil/utility/structure"
)

func Strip(s_json string) string {
	s_json = remove_comments(s_json)
	s_json = remove_configuration(s_json)
	s_json = remove_imports(s_json)

	return s_json
}

//
//
// Finds all comments and removes them
//
//
func remove_comments(s_json string) string {
	data_object := structure.Receive(s_json)
	regex := regexp.MustCompile(COMMENT)
	comments := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(comments) > 0 {
		for _, line := range comments[0] {
			data_object.File_gut = strings.Replace(data_object.File_gut, line, "", -1)
		}
	}

	return structure.Send(data_object)
}

//
//
// Removes the configuration section if it is found
//
//
func remove_configuration(s_json string) string {
	data_object := structure.Receive(s_json)
	regex := regexp.MustCompile(COMPILER_CONFIGURATION)
	comments := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(comments) > 0 {
		for _, line := range comments[0] {
			data_object.File_gut = strings.Replace(data_object.File_gut, line, "", -1)
		}
	}

	return structure.Send(data_object)
}

//
//
// Removes all imports from the structure
//
//
func remove_imports(s_json string) string {
	data_object := structure.Receive(s_json)
	regex := regexp.MustCompile(IMPORT)
	comments := regex.FindAllStringSubmatch(data_object.File_gut, -1)

	if len(comments) > 0 {
		for _, line := range comments[0] {
			data_object.File_gut = strings.Replace(data_object.File_gut, fmt.Sprintf("use %s", line), "", -1)
		}
	}

	return structure.Send(data_object)
}
