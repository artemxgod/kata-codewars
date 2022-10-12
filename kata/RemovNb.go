package kata

func RemovNb(m uint64) [][2]uint64 {
	sum := (m+1) * m / 2
	result := make([][2]uint64, 0)

	for a := sum / m; a <= m; a++ {
		b := (sum - a) / (a + 1);
		if a * b == sum - a - b && b <= m {
			result = append(result, [2]uint64{a, b})
		}
	}

	if len(result) == 0 {
		return nil
	}
	return result
}

// A friend of mine takes the sequence of all numbers from 1 to n (where n > 0).
// Within that sequence, he chooses two numbers, a and b.
// He says that the product of a and b should be equal to the sum of all numbers 
//in the sequence, excluding a and b.
// Given a number n, could you tell me the numbers he excluded from the sequence?
