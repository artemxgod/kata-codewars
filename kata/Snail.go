package kata


// `import "fmt"

func Snail(snaipMap [][]int) []int {
	if len(snaipMap) == 0 {
		return []int{}
	}
	if len(snaipMap[0]) == 0 {
		return []int{}
	}

	size, max := len(snaipMap), len(snaipMap)
	res := make([]int, 0)

	for idx := 0; idx < size; idx++ {
		if idx >= max { break }
		res = append(res, Clock(max, idx, snaipMap)...)	
		max--
	}
	return res
}

func Clock(max, min int, snaip [][]int) []int {
	idx, jdx := min, min
	res := make([]int, 0)
	for ; jdx < max; jdx++ {
		res = append(res, snaip[idx][jdx])
	}
	jdx--; idx++
	for ; idx < max; idx++ {
		res = append(res, snaip[idx][jdx])
	}
	idx--; jdx--
	for ; jdx >= min; jdx-- {
		res = append(res, snaip[idx][jdx])
	}
	jdx++; idx--
	for ; idx > min; idx-- {
		res = append(res, snaip[idx][jdx])
	}
	return res
}

// Snail Sort
// Given an n x n array, return the array elements arranged from outermost elements to the middle element, traveling clockwise.