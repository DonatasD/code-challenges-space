package main

import "fmt"

func strStr(haystack string, needle string) int {
	res := -1
	l := len(needle)
	for i := 0; i+l <= len(haystack); i++ {
		if haystack[i:i+l] == needle {
			return i
		}
	}
	return res
}

func main() {
	inputs := []string{"sadbutsad", "leetcode"}
	params := []string{"sad", "leeto"}
	outputs := []int{0, -1}

	for i := range inputs {

		result := strStr(inputs[i], params[i])
		if result == outputs[i] {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
