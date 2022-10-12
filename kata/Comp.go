package kata

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil || len(array1) != len(array2) {
		return false
	}
	for _, elem_1 := range array1 {
		check := false
		for idx, elem_2 := range array2 {
			if elem_2 == elem_1 * elem_1 {
				array2[idx] = 0
				check = true
				break
			}
		}
		if !check {
			return false
		}
	}
	return true
}

// TASK
// Compare 2 arrays. If second array contains all elements from first array in square, return true
// order does not matter