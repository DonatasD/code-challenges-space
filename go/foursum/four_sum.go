package main

import (
	"fmt"
	"reflect"
	"slices"
)

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	slices.Sort(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for l, r := j+1, len(nums)-1; l < r; {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					for l < r && nums[l] == nums[l-1] {
						l++
					}
				}
			}
		}
	}
	return res
}

func main() {
	inputs := [][]int{
		{1, 0, -1, 0, -2, 2},
		{2, 2, 2, 2, 2},
		{-3, -1, 0, 2, 4, 5},
	}
	params := []int{0, 8, 0}
	outputs := [][][]int{
		{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		{{2, 2, 2, 2}},
		{{-3, -1, 0, 4}},
	}
	for i := 0; i < len(inputs); i++ {
		result := fourSum(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
