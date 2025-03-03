package main

import (
	"fmt"
	"reflect"
)

func search(nums []int, target int) int {
	for l, r := 0, len(nums)-1; l <= r; {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}
		if nums[m] > target {
			r = m - 1
		}
		if nums[m] < target {
			l = m + 1
		}
	}
	return -1
}

func main() {
	inputs := [][]int{{-1, 0, 3, 5, 9, 12}, {-1, 0, 3, 5, 9, 12}}
	params := []int{9, 2}
	outputs := []int{4, -1}
	for i := 0; i < len(inputs); i++ {

		result := search(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
