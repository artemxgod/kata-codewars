package kata

func Parse(data string) []int {
	result := make([]int, 0)
	cur := 0
	for _, sym := range data {
		switch sym {
		case 'i':
			cur++
		case 'd':
			cur--
		case 's':
			cur *= cur
		case 'o':
			result = append(result, cur)
		}
	}

	return result
}
