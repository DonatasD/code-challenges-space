package main

import (
	"fmt"
	"math"
	"slices"
)

func threeSumClosest(nums []int, target int) int {
	res, bestDif := 0, math.MaxInt
	slices.Sort(nums)

	for i, v := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l, r := i+1, len(nums)-1; l < r; {
			sum := v + nums[l] + nums[r]
			dif := Abs(sum - target)
			if dif < bestDif {
				bestDif = dif
				res = sum
			}

			if sum == target {
				return target
			} else if sum > target {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func main() {
	inputs := [][]int{{-1, 2, 1, -4}, {0, 0, 0}}
	params := []int{1, 1}
	outputs := []int{2, 0}
	for i := 0; i < len(inputs); i++ {

		result := threeSumClosest(inputs[i], params[i])
		if result == outputs[i] {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
