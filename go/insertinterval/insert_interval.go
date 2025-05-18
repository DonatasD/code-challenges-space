package main

import (
	"fmt"
	"reflect"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	i := 0
	n := len(intervals)

	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	result = append(result, newInterval)

	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func main() {
	inputs := [][][]int{{{1, 3}, {6, 9}}, {{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}}
	params1 := [][]int{{2, 5}, {4, 8}}
	outputs := [][][]int{{{1, 5}, {6, 9}}, {{1, 2}, {3, 10}, {12, 16}}}
	for i := 0; i < len(inputs); i++ {

		result := insert(inputs[i], params1[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
