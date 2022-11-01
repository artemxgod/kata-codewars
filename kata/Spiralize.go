package kata

import "fmt"

// not my solution
func Spiralize(size int) [][]int {
  arr := make([][]int, size)
  for i := range arr { arr[i] = make([]int, size) }
  x,y,dx,dy,l := -2,0,1,0,size+1
  for l > 0 {
    for i := 0 ; i < l; i++ {
      x += dx
      y += dy
      if x >= 0 { arr[y][x] = 1 }
      if l == 1 { return arr }
    }
    if dy == 0 { l -= 2 }
    dx,dy = -dy,dx
	fmt.Println(y, x, dx, dy, l)
  }
  return arr
}
// func Spiralize(size int) [][]int {
// 	if size < 5 {
// 		return [][]int{}
// 	}

// 	workMatrix := make([][]int, size)
// 	for idx := range workMatrix { workMatrix[idx] = make([]int, size)}

// 	min := 0
// 	minj := 0
// 	br := false
// 	for idx := 0; idx < size; idx++ {
// 		max := size - idx
// 		workMatrix, br = fill(workMatrix, max, min, minj, size)
// 		min += 2
// 		minj = min - 1
// 		if br {
// 			break
// 		}
// 	}

// 	return workMatrix
// }

// func fill(matrix [][]int, max, min, minj, size int) ([][]int, bool) {
// 	if matrix[min][minj] == 1 {
// 		return matrix, true
// 	}
// 	idx, jdx := min, minj
// 	for ; jdx < max; jdx++ {
// 		if matrix[idx+1][jdx] == 1 {
// 			return matrix, true
// 		}
// 		matrix[idx][jdx] = 1
// 		if max < size && matrix[idx][jdx+1] == 1 {
// 			matrix[idx][jdx] = 0
// 			break
// 		}
// 	}
// 	jdx--; idx++
// 	for ; idx < max; idx++ {
// 		matrix[idx][jdx] = 1
// 		if max < size && matrix[idx+1][jdx] == 1 {
// 			matrix[idx][jdx] = 0
// 			// if idx == min + 2 {
// 			// 	return matrix, true
// 			// }
// 			break
// 		}
// 	}
// 	idx--; jdx--
// 	for ; jdx >= min; jdx-- {
// 		if idx != size - 1 {
// 			if matrix[idx-1][jdx] == 1 {
// 				return matrix, true
// 			}
// 		}
// 		matrix[idx][jdx] = 1
// 		if max < size && matrix[idx][jdx-1] == 1 {
// 			matrix[idx][jdx] = 0
// 			if idx == 6 && jdx == 5 {
// 				fmt.Println(matrix[idx][jdx])
// 			}
// 			break
// 		}
// 	}
// 	jdx++; idx--
// 	for ; idx > min; idx-- {
// 		matrix[idx][jdx] = 1
// 		if idx == min+1 && matrix[idx-1][jdx] == 1 {
// 			matrix[idx][jdx] = 0
// 			break
// 		}
// 	}

// 	return matrix, false
// }

func Print(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Print(matrix[i][j], " | ")
		}
		fmt.Println("")
	}
}

// Your task, is to create a NxN spiral with a given size.

// For example, spiral with size 5 should look like this:

// 00000
// ....0
// 000.0
// 0...0
// 00000