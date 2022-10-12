package kata

import (
	"strconv"
	"strings"
)

func Solution(list []int) string {
	var result string

	for idx := 0; idx < len(list); idx++ {
		tidx, tres := find_sec(list, idx)
		if tres {
			result += strconv.Itoa(list[idx]) + "-" + strconv.Itoa(list[tidx]) + ","
			idx = tidx
		} else {
			result += strconv.Itoa(list[idx]) + ","
		}
	}
	
	return strings.TrimSuffix(result, ",")
}

func find_sec(list []int, idx int) (int, bool) {
	rem := idx

	for idx+1 != len(list) {
		if list[idx+1]-1 == list[idx] {
			idx++
		} else {
			break
		}
	}

	if offset := idx - rem; offset < 2 {
		return 0, false
	} else {
		return idx, true
	}
}

//A format for expressing an ordered list 
//of integers is to use a comma separated 
//list of either
//individual integers
// or a range of integers denoted by 
//the starting integer separated 
//from the end integer in the range 
//by a dash, '-'. The range includes all integers 
//in the interval including both endpoints. 
//It is not considered a range unless 
//it spans at least 3 numbers. For example "12,13,15-17"

//Complete the solution so that it takes a list of integers 
//in increasing order and returns a correctly formatted string in the range format.