package main

import (
	"fmt"
	"reflect"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func main() {
	inputs := [][]int{{-2, 1, -3, 4, -1, 2, 1, -5, 4}, {1}, {5, 4, -1, 7, 8}}
	outputs := []int{6, 1, 23}
	for i := 0; i < len(inputs); i++ {
		result := maxSubArray(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, inputs[i])
		}
	}
}
