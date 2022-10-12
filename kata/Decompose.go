package kata

func Decompose(n int64) []int64 {
	return req(n*n, n)
}

func req(sum, idx int64) []int64 {
	if sum < 0 {
		return nil
	}
	if sum == 0 {
		return []int64{}
	}

	for jdx := idx -1; jdx > 0; jdx-- {
		sub := req(sum - jdx*jdx, jdx)
		if sub != nil {
			sub = append(sub, jdx)
			return sub
		}
	}
	return nil
}

// TASK:
// Given a positive integral number n,
// return a strictly increasing sequence
// (list/array/string depending on the language) of numbers,
// so that the sum of the squares is equal to n².

// If there are multiple solutions (and there will be),
// return as far as possible the result with the largest possible values: