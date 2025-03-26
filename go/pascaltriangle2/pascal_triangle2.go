package main

import (
	"fmt"
	"reflect"
)

func getRow(rowIndex int) []int {
	cur, prev := make([]int, rowIndex+1), make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				cur[j] = 1
			} else {
				cur[j] = prev[j-1] + prev[j]
			}
		}
		copy(prev, cur)
	}
	return cur
}

func main() {
	inputs := []int{
		3,
		2,
		0,
		1,
	}

	outputs := [][]int{
		{1, 3, 3, 1},
		{1, 2, 1},
		{1},
		{1, 1},
	}
	for i := 0; i < len(inputs); i++ {
		result := getRow(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
