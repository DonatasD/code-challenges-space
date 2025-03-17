package main

import (
	"fmt"
	"reflect"
)

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	perm := []int{}
	countMap := map[int]int{}
	for _, n := range nums {
		countMap[n] += 1
	}
	var dfs func()
	dfs = func() {
		if len(perm) == len(nums) {
			permCopy := append([]int{}, perm...)
			res = append(res, permCopy)
			return
		}

		for key, val := range countMap {
			if val > 0 {
				perm = append(perm, key)
				countMap[key] -= 1
				dfs()
				countMap[key] += 1
				perm = perm[:len(perm)-1]
			}
		}
	}
	dfs()
	return res
}

func main() {
	inputs := [][]int{
		{1, 2, 3},
		{1, 1, 2},
		{1},
	}
	outputs := [][][]int{
		{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		{{1}},
	}
	for i := 0; i < len(inputs); i++ {

		result := permuteUnique(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}

	}
}
