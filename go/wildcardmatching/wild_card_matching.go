package main

import (
	"fmt"
	"reflect"
)

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 1; j < n+1; j++ {
		if p[j-1:j] != "*" {
			break
		}
		dp[0][j] = true
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if p[j-1:j] == s[i-1:i] || p[j-1:j] == "?" {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1:j] == "*" {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

func main() {
	inputs := [][]string{
		{"adceb", "*a*b"},
		{"aa", "a"},
		{"aa", "*"},
		{"ab", "?a"},
	}
	outputs := []bool{
		true,
		false,
		true,
		false,
	}

	for i := 0; i < len(inputs); i++ {

		result := isMatch(inputs[i][0], inputs[i][1])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
