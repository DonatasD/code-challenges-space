package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	res := 0
	for x != 0 {
		digit := x % 10
		x = x / 10
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && digit >= math.MaxInt32%10) {
			return 0
		}
		if res < math.MinInt32/10 || (res == math.MinInt32/10 && digit <= math.MinInt32%10) {
			return 0
		}
		res = (res * 10) + digit
	}
	return res
}

func main() {
	inputs := []int{123, -123, 120, math.MaxInt32, math.MinInt32}
	outputs := []int{321, -321, 21, 0, 0}

	for i, v := range inputs {

		result := reverse(v)
		if result == outputs[i] {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
