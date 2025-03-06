package main

import (
	"fmt"
	"reflect"
)

func majorityElement(nums []int) []int {
	countMap := make(map[int]int)

	for _, v := range nums {
		existingCount, exists := countMap[v]
		if exists {
			countMap[v] = existingCount + 1
		} else {
			countMap[v] = 1
		}
		if len(countMap) <= 2 {
			continue
		}

		newCount := make(map[int]int)
		for k, c := range countMap {
			if c > 1 {
				newCount[k] = c - 1
			}
		}
		countMap = newCount
	}

	var res []int
	for k := range countMap {
		count := 0
		for _, val := range nums {
			if val == k {
				count++
			}
		}
		if count > len(nums)/3 {
			res = append(res, k)
		}
	}
	return res
}

func main() {
	inputs := [][]int{
		{3, 2, 3},
		{1},
		{1, 2},
		{2, 2, 1, 3},
	}
	outputs := [][]int{
		{3},
		{1},
		{1, 2},
		{2},
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
