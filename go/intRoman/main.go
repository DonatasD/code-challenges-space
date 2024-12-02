package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	decimal := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""
	for i := 0; i < len(decimal) && num > 0; i++ {
		n := num / decimal[i]
		result += strings.Repeat(roman[i], n)
		num -= n * decimal[i]
	}
	return result
}

func main() {
	inputs := []int{3749, 58, 1994}
	outputs := []string{"MMMDCCXLIX", "LVIII", "MCMXCIV"}

	for i, v := range inputs {

		result := intToRoman(v)
		if result == outputs[i] {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
