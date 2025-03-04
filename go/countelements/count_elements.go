package main

import (
	"fmt"
	"reflect"
)

func countElements(nums []int) int {
	minV, maxV := nums[0], nums[0]
	res, curMin, curMax := 0, 1, 1
	for _, v := range nums[1:] {
		if v > minV && v < maxV {
			res++
		}
		if v == maxV {
			curMax++
		}
		if v == minV {
			curMin++
		}
		if v < minV {
			if minV != nums[0] || maxV != nums[0] {
				res += curMin
				curMin = 1
			} else {
				curMin = 1
			}
			minV = v
		} else if v > maxV {
			if minV != nums[0] || maxV != nums[0] {
				res += curMax
				curMax = 1
			} else {
				curMax = 1
			}
			maxV = v
		}
	}
	return res
}

func main() {
	inputs := [][]int{
		{11, 7, 2, 15},
		{-3, 3, 3, 90},
	}
	outputs := []int{
		2,
		2,
	}
	for i := 0; i < len(inputs); i++ {

		result := countElements(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
