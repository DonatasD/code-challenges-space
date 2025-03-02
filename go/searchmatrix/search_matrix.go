package main

import (
	"fmt"
	"reflect"
)

func searchMatrix(matrix [][]int, target int) bool {
	x := -1
	for i := 0; i < len(matrix); i++ {
		if target == matrix[i][0] || target == matrix[i][len(matrix[i])-1] {
			return true
		}
		if target > matrix[i][0] && target < matrix[i][len(matrix[i])-1] {
			x = i
			break
		}
	}
	if x == -1 {
		return false
	}

	for l, r := 1, len(matrix[x])-2; l <= r; {
		m := (l + r) / 2
		if matrix[x][m] == target {
			return true
		}
		if matrix[x][m] > target {
			r = m - 1
		}
		if matrix[x][m] < target {
			l = m + 1
		}
	}
	return false
}

func main() {
	inputs := [][][]int{
		{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
		{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
		{{1, 3, 5}},
	}
	params := []int{
		3,
		13,
		3,
	}
	outputs := []bool{
		true,
		false,
		true,
	}
	for i := 0; i < len(inputs); i++ {

		result := searchMatrix(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
