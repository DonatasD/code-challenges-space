package main

import (
	"fmt"
	"reflect"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int

	var backtrack func(pos int, cur []int, target int)
	backtrack = func(pos int, cur []int, target int) {
		if target == 0 {
			res = append(res, append([]int{}, cur...))
			return
		}
		if target <= 0 {
			return
		}

		prev := -1
		for i := pos; i < len(candidates); i++ {
			if candidates[i] == prev {
				continue
			}
			cur = append(cur, candidates[i])
			backtrack(i+1, cur, target-candidates[i])
			cur = cur[:len(cur)-1]
			prev = candidates[i]
		}
	}

	backtrack(0, []int{}, target)
	return res
}

func main() {
	inputs := [][]int{{10, 1, 2, 7, 6, 1, 5}, {2, 5, 2, 1, 2}}
	params := []int{8, 5}
	outputs := [][][]int{{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}, {{1, 2, 2}, {5}}}
	for i := 0; i < len(inputs); i++ {

		result := combinationSum2(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
