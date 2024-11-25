package main

import "fmt"

func longestPalindrome(s string) string {
	array := []rune(s)
	resLen, resL, resR := 0, 0, 0
	for i := range array {
		l, r := i, i
		for l >= 0 && r < len(array) && array[l] == array[r] {
			if (r - l + 1) > resLen {
				resL = l
				resR = r + 1
				resLen = r - l + 1
			}
			l -= 1
			r += 1
		}
		l, r = i, i+1
		for l >= 0 && r < len(array) && array[l] == array[r] {
			if (r - l + 1) > resLen {
				resL = l
				resR = r + 1
				resLen = r - l + 1
			}
			l -= 1
			r += 1
		}
	}
	return string(array[resL:resR])
}

func main() {
	inputs := []string{"babad", "cbbd"}
	outputs := []string{"bab", "bb"}

	for i, v := range inputs {

		result := longestPalindrome(v)
		if result == outputs[i] {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v\n", i, result)
		}
	}
}
