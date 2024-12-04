package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	arr := []rune(s)
	var res []rune
	for r := range numRows {
		inc := 2 * (numRows - 1)
		for i := r; i < len(arr); i += inc {
			res = append(res, arr[i])
			if r > 0 && r < numRows-1 && i+inc-2*r < len(arr) {
				res = append(res, arr[i+inc-2*r])
			}
		}
	}
	return string(res)
}

func main() {
	inputs := []string{"PAYPALISHIRING", "PAYPALISHIRING", "PAYPALISHIRING"}
	params := []int{3, 4, 1}
	outputs := []string{"PAHNAPLSIIGYIR", "PINALSIGYAHRPI", "PAYPALISHIRING"}

	for i, v := range inputs {

		result := convert(v, params[i])
		if result == outputs[i] {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
