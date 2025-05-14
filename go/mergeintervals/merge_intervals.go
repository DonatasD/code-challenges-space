package main

import (
	"fmt"
	"reflect"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := result[len(result)-1]
		if current[0] <= lastMerged[1] {
			lastMerged[1] = max(lastMerged[1], current[1])
		} else {
			result = append(result, current)
		}
	}
	return result
}

func main() {
	inputs := [][][]int{{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, {{1, 4}, {4, 5}}}
	outputs := [][][]int{{{1, 6}, {8, 10}, {15, 18}}, {{1, 5}}}
	for i := 0; i < len(inputs); i++ {

		result := merge(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
