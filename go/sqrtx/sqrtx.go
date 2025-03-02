package main

import (
	"fmt"
	"reflect"
)

func mySqrt(x int) int {
	var search func(l int, r int, target int) int
	search = func(l int, r int, target int) int {
		mid := (l + r) / 2
		if r-l <= 1 {
			if (r * r) <= target {
				return r
			}
			return l
		}
		if mid*mid == target {
			return mid
		}
		if mid*mid > target {
			return search(l, mid, target)
		}
		return search(mid, r, target)
	}

	return search(0, x, x)
}

func main() {
	inputs := []int{4, 8, 2}
	outputs := []int{2, 2, 1}
	for i := 0; i < len(inputs); i++ {

		result := mySqrt(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
