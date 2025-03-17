package main

import (
	"fmt"
	"reflect"
)

func rotate(matrix [][]int) {
	l, r := 0, len(matrix)-1

	for l < r {
		for i := 0; i < r-l; i++ {
			top, bottom := l, r
			topLeft := matrix[top][l+i]
			matrix[top][l+i] = matrix[bottom-i][l]
			matrix[bottom-i][l] = matrix[bottom][r-i]
			matrix[bottom][r-i] = matrix[top+i][r]
			matrix[top+i][r] = topLeft
		}
		r -= 1
		l += 1
	}
}

func main() {
	inputs := [][][]int{
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
	}
	outputs := [][][]int{
		{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
	}
	for i := 0; i < len(inputs); i++ {

		rotate(inputs[i])
		if reflect.DeepEqual(inputs[i], outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, inputs[i], outputs[i])
		}

	}
}
