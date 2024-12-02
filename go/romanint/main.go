package main

import (
	"fmt"
)

func romanToInt(s string) int {
	roman := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	res := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && roman[s[i:i+1]] < roman[s[i+1:i+2]] {
			res -= roman[s[i:i+1]]
		} else {
			res += roman[s[i:i+1]]
		}
	}
	return res
}

func main() {
	inputs := []string{"MMMDCCXLIX", "LVIII", "MCMXCIV", "DCXXI"}
	outputs := []int{3749, 58, 1994, 621}

	for i, v := range inputs {

		result := romanToInt(v)
		if result == outputs[i] {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
