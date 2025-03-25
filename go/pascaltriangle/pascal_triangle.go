package main

import (
	"fmt"
	"reflect"
)

func generate(numRows int) [][]int {
	var res [][]int
	for i := 0; i < numRows; i++ {
		var cur []int
		for j := 0; j <= i; j++ {
			if j == 0 || i == j {
				cur = append(cur, 1)
			} else {
				cur = append(cur, res[i-1][j-1]+res[i-1][j])
			}
		}
		res = append(res, cur)
	}
	return res
}

func main() {
	inputs := []int{
		5,
		1,
	}

	outputs := [][][]int{
		{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		{{1}},
	}
	for i := 0; i < len(inputs); i++ {
		result := generate(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
