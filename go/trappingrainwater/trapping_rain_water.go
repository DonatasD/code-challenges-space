package main

import (
	"fmt"
	"reflect"
)

func trap(height []int) int {
	var maxL = make([]int, len(height))
	var maxR = make([]int, len(height))

	for i := 0; i < len(height); i++ {
		val := 0
		if i > 0 {
			val = max(maxL[i-1], height[i-1])
		}
		maxL[i] = val
	}

	for i := len(height) - 1; i >= 0; i-- {
		val := 0
		if i < len(height)-1 {
			val = max(maxR[i+1], height[i+1])
		}
		maxR[i] = val
	}

	sum := 0
	for i := 0; i < len(height); i++ {
		val := min(maxL[i], maxR[i]) - height[i]
		if val > 0 {
			sum += val
		}
	}
	return sum
}

func trapOptimal(height []int) int {
	if len(height) < 3 {
		return 0
	}
	l, r := 0, len(height)-1
	lMax, rMax := height[0], height[len(height)-1]
	res := 0

	for l < r {
		if lMax < rMax {
			l += 1
			lMax = max(lMax, height[l])
			res += lMax - height[l]
		} else {
			r -= 1
			rMax = max(rMax, height[r])
			res += rMax - height[r]
		}
	}
	return res
}

func main() {
	inputs := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{4, 2, 0, 3, 2, 5},
	}
	outputs := []int{
		6,
		9,
	}
	for i := 0; i < len(inputs); i++ {
		result := trap(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}

		resultOptimal := trapOptimal(inputs[i])
		if reflect.DeepEqual(resultOptimal, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
