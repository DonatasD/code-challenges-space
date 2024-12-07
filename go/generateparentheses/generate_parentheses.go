package main

import (
	"fmt"
	"reflect"
)

func generateParenthesis(n int) []string {
	var res []string
	var backtrack func(cur string, open int, closed int)
	backtrack = func(cur string, open int, closed int) {
		if closed == n && open == n {
			res = append(res, cur)
			return
		}
		if open < n {
			backtrack(cur+"(", open+1, closed)
		}
		if closed < open {
			backtrack(cur+")", open, closed+1)
		}
	}
	backtrack("", 0, 0)
	return res
}

func main() {
	inputs := []int{3, 1}
	outputs := [][]string{{"((()))", "(()())", "(())()", "()(())", "()()()"}, {"()"}}
	for i := 0; i < len(inputs); i++ {
		result := generateParenthesis(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
