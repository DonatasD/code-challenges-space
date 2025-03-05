package main

import (
	"fmt"
	"reflect"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func firstMissingPositive(nums []int) int {
	for i := range nums {
		if nums[i] < 0 {
			nums[i] = 0
		}
	}

	for i := range nums {
		v := abs(nums[i])
		if 1 <= v && v <= len(nums) {
			if nums[v-1] > 0 {
				nums[v-1] *= -1
			} else if nums[v-1] == 0 {
				nums[v-1] = -(len(nums) + 1)
			}
		}
	}

	for i := 1; i < len(nums)+1; i++ {
		if nums[i-1] >= 0 {
			return i
		}

	}
	return len(nums) + 1
}

func main() {
	inputs := [][]int{
		{1, 2, 0},
		{3, 4, -1, 1},
		{7, 8, 9, 11, 12},
		{1},
	}
	outputs := []int{
		3,
		2,
		1,
		2,
	}
	for i := 0; i < len(inputs); i++ {

		result := firstMissingPositive(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}

	}
}
