package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func reverse(input string) string {
	n := 0
	runeVal := make([]rune, len(input))

	for _, r := range input {
		runeVal[n] = r
		n++
	}

	runeVal = runeVal[0:n]
	for i := 0; i < n/2; i++ {
		runeVal[i], runeVal[n-1-i] = runeVal[n-1-i], runeVal[i]
	}
	return string(runeVal)
}

func addBinary(a string, b string) string {
	res := ""
	carry := 0
	a, b = reverse(a), reverse(b)

	for i := 0; i < max(len(a), len(b)); i++ {
		digitA, digitB := 0, 0
		if i < len(a) {
			digitA = int(a[i] - '0')
		}

		if i < len(b) {
			digitB = int(b[i] - '0')
		}

		total := digitA + digitB + carry
		char := strconv.Itoa(total % 2)
		res = char + res
		carry = total / 2
	}
	if carry > 0 {
		res = "1" + res
	}
	return res
}

func main() {
	inputs := [][]string{
		{"11", "1"},
		{"1010", "1011"},
	}
	outputs := []string{"100", "10101"}
	for i := 0; i < len(inputs); i++ {

		result := addBinary(inputs[i][0], inputs[i][1])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}

	}
}
