package main

func find(array []string, value string) bool {
	for _, entry := range array {
		if entry == value {
			return true
		}
	}
	return false
}
