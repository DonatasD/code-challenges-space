package main

import (
	"fmt"
	"reflect"
	"sort"
)

func triangleType(nums []int) string {
	sort.Ints(nums)
	if nums[0]+nums[1] <= nums[2] {
		return "none"
	}
	if nums[0] == nums[1] && nums[1] == nums[2] {
		return "equilateral"
	}
	if nums[0] == nums[1] || nums[1] == nums[2] || nums[0] == nums[2] {
		return "isosceles"
	}
	return "scalene"
}

func main() {
	inputs := [][]int{{3, 3, 3}, {3, 4, 5}}
	outputs := []string{"equilateral", "scalene"}
	for i := 0; i < len(inputs); i++ {
		result := triangleType(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
