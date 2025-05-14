package main

import (
	"fmt"
	"reflect"
)

func canJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
	}
	return true
}

func main() {
	inputs := [][]int{{2, 3, 1, 1, 4}, {3, 2, 1, 0, 4}}
	outputs := []bool{true, false}
	for i := 0; i < len(inputs); i++ {

		result := canJump(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
