package main

import "fmt"

func isMatch(s string, p string) bool {

	cache := make(map[string]bool)
	var dfs func(i int, j int) bool
	dfs = func(i int, j int) bool {
		key := fmt.Sprintf("%d %d", i, j)
		val, exists := cache[key]
		if exists {
			return val
		}
		if i >= len(s) && j >= len(p) {
			return true
		}
		if j >= len(p) {
			return false
		}
		match := i < len(s) && (s[i] == p[j] || string(p[j]) == ".")
		if (j+1) < len(p) && string(p[j+1]) == "*" {
			cache[key] = dfs(i, j+2) || (match && dfs(i+1, j))
			return cache[key]
		}

		if match {
			cache[key] = dfs(i+1, j+1)
			return cache[key]
		}

		return false
	}
	return dfs(0, 0)
}

func main() {
	inputs := []string{"aa", "aa", "ab"}
	params := []string{"a", "a*", ".*"}
	outputs := []bool{false, true, true}

	for i, v := range inputs {

		result := isMatch(v, params[i])
		if result == outputs[i] {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
