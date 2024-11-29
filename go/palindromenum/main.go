package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	m := len(s) / 2
	if len(s)%2 == 0 {
		return s[:m] == reverse(s[m:])
	}
	return s[:m] == reverse(s[m+1:])
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	inputs := []int{13231, -121, 10, 123}
	outputs := []bool{true, false, false, false}

	for i, v := range inputs {

		result := isPalindrome(v)
		if result == outputs[i] {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
