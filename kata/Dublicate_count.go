package kata

import "strings"

func Duplicate_count(s1 string) int {
	result := 0
	s1 = strings.ToLower(s1)
	for _, sym := range s1 {
		if sym == '_' {
			continue
		}
		if strings.Count(s1, string(sym)) > 1 {
			s1 = strings.ReplaceAll(s1, string(sym), "_")
			result++
		}
	}
	return result
}


// TASK:
// Write a function that will return the count of distinct case-insensitive alphabetic characters and numeric digits 
// that occur more than once in the input string. 
// The input string can be assumed to contain only alphabets (both uppercase and lowercase) and numeric digits.