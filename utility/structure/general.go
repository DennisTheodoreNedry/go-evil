package structure

import (
	"encoding/base64"
	"encoding/json"

	compilevar "github.com/s9rA16Bf4/go-evil/utility/structure/compile_var"
	evilarray "github.com/s9rA16Bf4/go-evil/utility/structure/evil_array"
	evil_json "github.com/s9rA16Bf4/go-evil/utility/structure/json"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Creates a json object and returns it
func Create_json_object() evil_json.Json_t {
	var new_json evil_json.Json_t
	new_json.Malware_src_file = "malware.go"
	new_json.Malware_gut = append(new_json.Malware_gut, "package main")

	// Default settings for all the window instances
	new_json.Width = 800
	new_json.Height = 600
	new_json.Title = "me_no_virus"

	// Used when we are doing a 1:1 mapping of our js function to an evil one
	new_json.Bind_gut = make(map[string]string)

	new_json.Var_max = 5

	for id := 0; id < new_json.Var_max; id++ {
		var new_var compilevar.Compile_var_t
		new_var.Set_value("NULL") // Default value
		new_json.Comp_var = append(new_json.Comp_var, new_var)
	}

	new_json.Set_alphabet("0,1,2,3,4,5,6,7,8,9,a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z,A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z,!,#,$,€,%,&,(,),*,+,,,-,.,/,:,;,<,=,>,?,@,[,],^,_,`,{,|,},~, ,%t,%n,%r,%x0b,%x0c")

	return new_json
}

// Creates an evil array object and returns it
func Create_evil_object(arr_content string) evilarray.Evil_array_t {
	var new_arr evilarray.Evil_array_t
	new_arr.Set_length(0)
	new_arr.Set_gut([]string{})

	new_arr.Parse(arr_content) // Will populate the array with the provided gut

	return new_arr
}

// Serializes the json structure into a base64 string which is ready to be sent
func Send(object evil_json.Json_t) string {
	serial_json, err := json.Marshal(object)

	if err != nil {
		notify.Error(err.Error(), "json.Convert_to_json()")
	}

	return base64.StdEncoding.EncodeToString(serial_json)
}

// Used to convert the received serialized json structure into workable data
func Receive(object string) evil_json.Json_t {
	serialize_json, err := base64.StdEncoding.DecodeString(object)

	if err != nil {
		notify.Error(err.Error(), "json.Receive()")
	}

	var result evil_json.Json_t
	if err := json.Unmarshal(serialize_json, &result); err != nil {
		notify.Error(err.Error(), "json.Convert_to_Json_t()")
	}

	return result
}
