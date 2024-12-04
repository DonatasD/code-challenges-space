package main

import (
	"fmt"
	"reflect"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		n, exists := m[target-v]
		if exists {
			return []int{n, i}
		}
		m[v] = i
	}
	return []int{}
}

func main() {
	inputs := [][]int{{2, 7, 11, 15}, {3, 2, 4}, {3, 3}}
	params := []int{9, 6, 6}
	outputs := [][]int{{0, 1}, {1, 2}, {0, 1}}
	for i := 0; i < len(inputs); i++ {

		result := twoSum(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
