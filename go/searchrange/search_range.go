package main

import (
	"fmt"
	"reflect"
)

func searchRange(nums []int, target int) []int {
	var binSearch func(findLeft bool) int
	binSearch = func(findLeft bool) int {
		l, r := 0, len(nums)-1
		i := -1
		for l <= r {
			m := (l + r) / 2
			if target > nums[m] {
				l = m + 1
			} else if target < nums[m] {
				r = m - 1
			} else {
				i = m
				if findLeft {
					r = m - 1
				} else {
					l = m + 1
				}
			}
		}
		return i
	}
	left := binSearch(true)
	right := binSearch(false)
	return []int{left, right}
}

func main() {
	inputs := [][]int{{5, 7, 7, 8, 8, 10}, {5, 7, 7, 8, 8, 10}, {}}
	params := []int{8, 6, 0}
	outputs := [][]int{{3, 4}, {-1, -1}, {-1, -1}}
	for i := 0; i < len(inputs); i++ {
		result := searchRange(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
