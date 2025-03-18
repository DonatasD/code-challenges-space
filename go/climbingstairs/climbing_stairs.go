package main

import (
	"fmt"
	"reflect"
)

func climbStairs(n int) int {
	one, two := 1, 1

	for i := 0; i < n-1; i++ {
		one, two = two+one, one
	}

	return one
}

func main() {
	inputs := []int{2, 3}
	outputs := []int{2, 3}
	for i := 0; i < len(inputs); i++ {

		result := climbStairs(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
