package main

import (
	"fmt"
	"reflect"
)

func letterCombinations(digits string) []string {
	var res []string
	digitsMap := map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "xwyz"}

	var backtrack func(i int, cur string)
	backtrack = func(i int, cur string) {
		if len(cur) == len(digits) {
			res = append(res, cur)
			return
		}
		chars := digitsMap[digits[i:i+1]]
		for j := range chars {
			backtrack(i+1, cur+chars[j:j+1])
		}
	}
	if digits != "" {
		backtrack(0, "")
	}
	return res
}

func main() {
	inputs := []string{"23", "", "2"}
	outputs := [][]string{{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, {}, {"a", "b", "c"}}
	for i := 0; i < len(inputs); i++ {

		result := letterCombinations(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
