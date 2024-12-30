package main

import (
	"fmt"
	"reflect"
)

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2
		if target == nums[m] {
			return m
		}

		if nums[l] <= nums[m] {
			if target > nums[m] || target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if target < nums[m] || target > nums[r] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return -1
}

func main() {
	inputs := [][]int{{4, 5, 6, 7, 0, 1, 2}, {4, 5, 6, 7, 0, 1, 2}, {1}}
	params := []int{0, 3, 0}
	outputs := []int{4, -1, -1}
	for i := 0; i < len(inputs); i++ {
		result := search(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
