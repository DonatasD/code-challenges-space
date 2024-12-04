package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	res := ""
	for i, _ := range strs[0] {
		for _, v := range strs {
			if i == len(v) || strs[0][i:i+1] != v[i:i+1] {
				return res
			}
		}
		res += strs[0][i : i+1]
	}
	return res
}

func main() {
	inputs := [][]string{{"flower", "flow", "flight"}, {"dog", "racecar", "car"}}
	outputs := []string{"fl", ""}

	for i := 0; i < len(inputs); i++ {

		result := longestCommonPrefix(inputs[i])
		if result == outputs[i] {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
