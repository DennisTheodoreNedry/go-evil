package functions

type Func_t struct {
	Name        string   // Obfuscated function name or the real function name
	Func_type   string   // Which sort of type this function is
	Return_type string   // What kind of value can we expect to be returned
	Gut         []string // The contents of the function
}

type Go_func_t struct {
	Name           string   // Obfuscated function name or the real function name
	Func_type      string   // Which sort of type this function is, call/loop/boot
	Part_of_struct string   // This should contain the name of what struct this function is part of
	Return_type    string   // What kind of value can we expect to be returned
	Parameters     []string // Potential input needed
	Gut            []string // The contents of the function
}
