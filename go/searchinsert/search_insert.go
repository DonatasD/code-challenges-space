package main

import (
	"fmt"
	"reflect"
)

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (r + l) / 2
		if target == nums[m] {
			return m
		}

		if target > nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

func main() {
	inputs := [][]int{{1, 3, 5, 6}, {1, 3, 5, 6}, {1, 3, 5, 6}}
	params := []int{5, 2, 7}
	outputs := []int{2, 1, 4}
	for i := 0; i < len(inputs); i++ {
		result := searchInsert(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
