package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	decimal := []int{1000, 500, 100, 50, 10, 5, 1}
	result := ""
	roman := []string{"M", "D", "C", "L", "X", "V", "I"}
	for i, v := range decimal {
		n := num / v
		if n > 0 && n < 4 && num != 9 {
			result += strings.Repeat(roman[i], n)
			num -= n * v
		} else if num == 9 {
			result += "IX"
			num -= num
		} else if n > 3 {
			result += roman[i] + roman[i-1]
			num -= decimal[i-1] - decimal[i]
		}
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
