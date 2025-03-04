package main

import (
	"fmt"
	"reflect"
)

func missingNumber(nums []int) int {
	total := (len(nums) * (len(nums) + 1)) / 2
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return total - sum
}

func missingNumber1(nums []int) int {
	res := len(nums)
	for i, v := range nums {
		res += i - v
	}
	return res
}

func main() {
	inputs := [][]int{
		{3, 0, 1},
		{0, 1},
		{9, 6, 4, 2, 3, 5, 7, 0, 1},
	}
	outputs := []int{
		2,
		2,
		8,
	}
	for i := 0; i < len(inputs); i++ {

		result := missingNumber(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}

		result1 := missingNumber1(inputs[i])
		if reflect.DeepEqual(result1, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
