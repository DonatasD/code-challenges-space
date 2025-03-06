package main

import (
	"fmt"
	"reflect"
)

func majorityElement(nums []int) int {
	res, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			res = v
		}
		if v == res {
			count++
		} else {
			count--
		}
	}
	return res
}

func main() {
	inputs := [][]int{
		{3, 2, 3},
		{2, 2, 1, 1, 1, 2, 2},
		{3, 3, 4},
		{8, 8, 7, 7, 7},
	}
	outputs := []int{
		3,
		2,
		3,
		7,
	}
	for i := 0; i < len(inputs); i++ {

		result := majorityElement(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}

	}
}
