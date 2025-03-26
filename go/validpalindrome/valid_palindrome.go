package main

import (
	"fmt"
	"reflect"
	"unicode"
)

func isPalindrome(s string) bool {
	var isAlphaNum func(c rune) bool
	isAlphaNum = func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsDigit(c)
	}
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphaNum(rune(s[l])) {
			l++
		}
		for r > l && !isAlphaNum(rune(s[r])) {
			r--
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	inputs := []string{"A man, a plan, a canal: Panama", "race a car", " "}
	outputs := []bool{true, false, true}
	for i := 0; i < len(inputs); i++ {

		result := isPalindrome(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
