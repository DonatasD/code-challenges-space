package main

import (
	"fmt"
	"reflect"
	"slices"
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

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]int, len(num1)+len(num2))
	num1, num2 = reverse(num1), reverse(num2)

	for i, v1 := range num1 {
		for j, v2 := range num2 {
			digit := int(v1-'0') * int(v2-'0')
			res[i+j] += digit
			res[i+j+1] += res[i+j] / 10
			res[i+j] = res[i+j] % 10
		}
	}
	slices.Reverse(res)
	beg := 0
	for beg < len(res) && res[beg] == 0 {
		beg += 1
	}
	resString := ""
	for i := beg; i < len(res); i++ {
		resString += fmt.Sprintf("%d", res[i])
	}
	return resString
}

func main() {
	inputs := [][]string{
		{"2", "3"},
		{"123", "456"},
	}
	outputs := []string{
		"6",
		"56088",
	}
	for i := 0; i < len(inputs); i++ {

		result := multiply(inputs[i][0], inputs[i][1])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}

	}
}
