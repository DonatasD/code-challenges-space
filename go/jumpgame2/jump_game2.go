package main

import (
	"fmt"
	"reflect"
)

func jump(nums []int) int {
	res, l, r := 0, 0, 0
	for r < len(nums)-1 {
		farthest := 0
		for i := l; i < r+1; i++ {
			farthest = max(farthest, i+nums[i])
		}
		l = r + 1
		r = farthest
		res += 1
	}
	return res
}

func main() {
	inputs := [][]int{
		{2, 3, 1, 1, 4},
		{2, 3, 0, 1, 4},
		{1, 2, 3},
		{2, 1, 1, 1, 1},
	}
	outputs := []int{
		2,
		2,
		2,
		3,
	}

	for i := 0; i < len(inputs); i++ {

		result := jump(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}

	}
}
