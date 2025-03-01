package main

import (
	"fmt"
	"reflect"
)

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var dfs func(i int, cur []int, total int)
	dfs = func(i int, cur []int, total int) {
		if total == target {
			res = append(res, append([]int{}, cur...))
			return
		}
		if i >= len(candidates) || total > target {
			return
		}

		cur = append(cur, candidates[i])
		dfs(i, cur, total+candidates[i])
		cur = cur[:len(cur)-1]
		dfs(i+1, cur, total)
	}

	dfs(0, []int{}, 0)
	return res
}

func main() {
	inputs := [][]int{{2, 3, 6, 7}, {2, 3, 5}, {2}}
	params := []int{7, 8, 1}
	outputs := [][][]int{{{2, 2, 3}, {7}}, {{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, {}}
	for i := 0; i < len(inputs); i++ {

		result := combinationSum(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
