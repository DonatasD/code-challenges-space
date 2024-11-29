package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	arr := []rune(strings.TrimLeft(s, " "))
	if len(arr) == 0 {
		return 0
	}
	i := 0
	sign := 1
	cur := arr[i]

	if string(cur) == "+" {
		i++
	} else if string(cur) == "-" {
		i++
		sign = -1
	}

	res := 0
	for i < len(arr) {
		cur := arr[i]
		i++
		if !unicode.IsDigit(cur) {
			break
		}

		res = res*10 + int(cur-'0')
		if res*sign >= math.MaxInt32 {
			return math.MaxInt32
		}
		if res*sign <= math.MinInt32 {
			return math.MinInt32
		}
	}
	return res * sign
}

func main() {
	inputs := []string{"42", " -042", "1337c0d3", "0-1", "words and 987", "2147483648123", "-2147483648123"}
	outputs := []int{42, -42, 1337, 0, 0, math.MaxInt32, math.MinInt32}

	for i, v := range inputs {

		result := myAtoi(v)
		if result == outputs[i] {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
