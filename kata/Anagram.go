package kata

import "strings"

func Anagrams(word string, words []string) []string {
	result := make([]string, 0)

	for _, elem := range words {
		if len(elem) == len(word) {
			check := true
			for _, sym := range word {
				if strings.Count(word, string(sym)) != strings.Count(elem, string(sym)) {
					check = false
				}
			}
			if check {
				result = append(result, elem)
			}
		}	
	}
	return result
}

// TASK:
// return all anagrams for the word given as first param