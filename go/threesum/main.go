package main

import (
	"fmt"
	"reflect"
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	var res [][]int
	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}
		for l, r := i+1, len(nums)-1; l < r; {
			sum := v + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{v, nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return res
}

func main() {
	inputs := [][]int{{-1, 0, 1, 2, -1, -4}, {0, 1, 1}, {0, 0, 0}, {1, -1, -1, 0}}
	outputs := [][][]int{{{-1, -1, 2}, {-1, 0, 1}}, {}, {{0, 0, 0}}, {{-1, 0, 1}}}
	for i := 0; i < len(inputs); i++ {
		//for i, v := range inputs {

		result := threeSum(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
