package main

import (
	"fmt"
	"reflect"
	"sort"
)

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	valueMap := make(map[int]int)
	for _, item := range items1 {
		valueMap[item[0]] += item[1]
	}
	for _, item := range items2 {
		valueMap[item[0]] += item[1]
	}

	result := make([][]int, 0, len(valueMap))
	for value, weight := range valueMap {
		result = append(result, []int{value, weight})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

func main() {
	inputs := [][][][]int{
		{{{1, 1}, {4, 5}, {3, 8}}, {{3, 1}, {1, 5}}},
		{{{1, 1}, {3, 2}, {2, 3}}, {{2, 1}, {3, 2}, {1, 3}}},
		{{{1, 3}, {2, 2}}, {{7, 1}, {2, 2}, {1, 4}}},
	}
	outputs := [][][]int{
		{{1, 6}, {3, 9}, {4, 5}},
		{{1, 4}, {2, 4}, {3, 4}},
		{{1, 7}, {2, 4}, {7, 1}},
	}
	for i := 0; i < len(inputs); i++ {
		result := mergeSimilarItems(inputs[i][0], inputs[i][1])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
