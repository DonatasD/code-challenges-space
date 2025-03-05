package main

import (
	"fmt"
	"reflect"
)

func prependInt(x []int, y int) []int {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}

func plusOne(digits []int) []int {
	var prepend bool
	for i := len(digits) - 1; i >= 0; i-- {
		prepend = false
		if digits[i] == 9 {
			digits[i] = 0
			prepend = true
		} else {
			digits[i]++
			break
		}
	}
	if prepend {
		return prependInt(digits, 1)
	}
	return digits
}

func main() {
	inputs := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
	}
	outputs := [][]int{
		{1, 2, 4},
		{4, 3, 2, 2},
		{1, 0},
	}
	for i := 0; i < len(inputs); i++ {

		result := plusOne(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}

	}
}
