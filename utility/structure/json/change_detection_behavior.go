package json

// Sets the behavior tactic that the malware will undertaken
// if it detects that it's being launched with the help of a debugger
func (object *Json_t) Change_detection_behavior(tactic string) string {
	object.Debugger_behavior = tactic
	return ""
}
