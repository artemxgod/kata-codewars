package kata

import (
	"fmt"
	"math"
)

func SumOfDivided(lst []int) string {
	var result string

	abs_sort(lst)

	num := 2
	sum := 0
	flag := false
	for _, elem := range lst {
		if elem % 2 == 0 {
			sum += elem
			if !flag {
				flag = true
			}
		}
	}
	if flag {
		result += fmt.Sprintf("(%d %d)", num, sum)
	}
	for idx := 3; idx <= int(math.Abs(float64(lst[len(lst)-1]))); idx += 2 {
		num := 0
		sum := 0
		flag := false
		if !is_prime(idx) { continue }
		for _, elem := range lst {
			if elem % idx == 0 {
				sum += elem
				if !flag {
					flag = true
					num = idx
				}
			}
		}
		if flag {
			result += fmt.Sprintf("(%d %d)", num, sum)
		}
	}

	return result
}

func abs_sort(lst []int) []int{
	for i := 1; i < len(lst); i++ {
		j := i
		for j > 0 && math.Abs(float64(lst[j])) < math.Abs(float64(lst[j-1])) {
			lst[j], lst[j-1] = lst[j-1], lst[j]
			j--
		}
	}
	return lst
}

func is_prime(num int) bool {
	if num == 1 {
		return true
	}
	for i := 2; i <= num / 2; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}


// Given an array of positive or negative integers

//  I= [i1,..,in]

// you have to produce a sorted array P of the form

// [ [p, sum of all ij of I for which p is a prime factor (p positive) of ij] ...]

// P will be sorted by increasing order of the prime numbers. 
// The final result has to be given as a string in Java, C#, C, C++
// and as an array of arrays in other languages.