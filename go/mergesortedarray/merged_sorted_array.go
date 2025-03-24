package main

import (
	"fmt"
	"reflect"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

func main() {
	inputs := [][][]int{
		{{4, 5, 6, 0, 0, 0}, {1, 2, 3}},
		{{1, 2, 3, 0, 0, 0}, {2, 5, 6}},
		{{1}, {}},
		{{0}, {1}},
	}
	params := [][]int{
		{3, 3},
		{3, 3},
		{1, 0},
		{0, 1},
	}
	outputs := [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 2, 3, 5, 6},
		{1},
		{1},
	}
	for i := 0; i < len(inputs); i++ {
		merge(inputs[i][0], params[i][0], inputs[i][1], params[i][1])
		if reflect.DeepEqual(inputs[i][0], outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, inputs[i][0], outputs[i])
		}
	}
}
