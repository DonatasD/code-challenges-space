package main

import (
	"fmt"
	"reflect"
)

func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}

	return res
}

func main() {
	inputs := [][][]int{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, {{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}}
	outputs := [][]int{{1, 2, 3, 6, 9, 8, 7, 4, 5}, {1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}}
	for i := 0; i < len(inputs); i++ {

		result := spiralOrder(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
