package main

import (
	"fmt"
	"math"
	"reflect"
)

func longestValidParentheses(s string) int {
	l, r, res := 0, 0, 0

	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "(" {
			l++
		} else {
			r++
		}

		if l == r {
			res = int(math.Max(float64(res), float64(l+r)))
		} else if r > l {
			l, r = 0, 0
		}
	}
	l, r = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i:i+1] == "(" {
			l++
		} else {
			r++
		}

		if l == r {
			res = int(math.Max(float64(res), float64(l+r)))
		} else if l > r {
			l, r = 0, 0
		}
	}

	return res
}

func main() {
	inputs := []string{"(()", ")()())", ""}
	outputs := []int{2, 4, 0}
	for i := 0; i < len(inputs); i++ {
		result := longestValidParentheses(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
