package main

import (
	"fmt"
	"reflect"
)

func insert(array []int, element int, i int) []int {
	return append(array[:i], append([]int{element}, array[i:]...)...)
}

func permute(nums []int) [][]int {
	perms := [][]int{{}}
	for _, n := range nums {
		var newPerms [][]int
		for _, p := range perms {
			for i := 0; i <= len(p); i++ {
				pCopy := append([]int{}, p...)
				pCopy = insert(pCopy, n, i)
				newPerms = append(newPerms, pCopy)
			}
		}
		perms = newPerms
	}
	return perms
}

func main() {
	inputs := [][]int{
		{1, 2, 3},
		{0, 1},
		{1},
	}
	outputs := [][][]int{
		{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		{{0, 1}, {1, 0}},
		{{1}},
	}
	for i := 0; i < len(inputs); i++ {

		result := permute(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}

	}
}
