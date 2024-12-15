package main

import (
	"fmt"
	"reflect"
	"slices"
)

func nextPermutation(nums []int) {
	pivot := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			pivot = i - 1
			break
		}
	}

	if pivot == -1 {
		slices.Reverse(nums)
		return
	}

	swap := len(nums) - 1
	for nums[swap] <= nums[pivot] {
		swap--
	}

	nums[pivot], nums[swap] = nums[swap], nums[pivot]
	slices.Reverse(nums[pivot+1:])
}

func main() {
	inputs := [][]int{{1, 2, 3}, {3, 2, 1}, {1, 1, 5}, {2, 3, 1}}
	outputs := [][]int{{1, 3, 2}, {1, 2, 3}, {1, 5, 1}, {3, 1, 2}}
	for i := 0; i < len(inputs); i++ {
		nextPermutation(inputs[i])
		if reflect.DeepEqual(inputs[i], outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, inputs[i], outputs[i])
		}
	}
}
