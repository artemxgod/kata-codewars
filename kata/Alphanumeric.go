package kata

import "unicode"


func Alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, elem := range str {
		if !(unicode.IsDigit(elem) || unicode.IsLetter(elem)) {
			return false
		}
	}
	return true
}

// TASK:
// check if string is alphanumeric (contains only latin alphabet symbols and digits)