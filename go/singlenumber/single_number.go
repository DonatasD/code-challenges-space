package main

import (
	"fmt"
	"reflect"
)

func singleNumber(nums []int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i]
	}

	return result
}

func main() {
	inputs := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
		{1},
	}
	outputs := []int{1, 4, 1}
	for i := 0; i < len(inputs); i++ {

		result := singleNumber(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
