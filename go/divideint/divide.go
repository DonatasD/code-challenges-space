package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	res, sign := 0, 1
	if dividend < 0 {
		sign = -sign
		dividend = -dividend
	}

	if divisor < 0 {
		sign = -sign
		divisor = -divisor
	}

	for dividend >= divisor {
		tmp := divisor
		mul := 1
		for dividend >= tmp {
			dividend -= tmp
			res += mul
			mul += mul
			tmp += tmp
		}
	}

	if sign < 0 {
		res = -res
	}

	if res < math.MinInt32 {
		return math.MinInt32
	}

	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	return res
}

func main() {
	inputs := []int{10, 7, -2147483648}
	params := []int{3, -3, -3}
	outputs := []int{3, -2, 715827882}

	for i := range inputs {
		result := divide(inputs[i], params[i])
		if result == outputs[i] {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
