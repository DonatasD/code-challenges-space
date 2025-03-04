package main

import (
	"fmt"
	"reflect"
)

func mergeAlternately(word1 string, word2 string) string {
	res := ""
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			res += word1[i : i+1]
		}
		if i < len(word2) {
			res += word2[i : i+1]
		}
	}
	return res
}

func main() {
	inputs := []string{
		"abc",
		"ab",
		"abcd",
	}
	params := []string{
		"pqr",
		"pqrs",
		"pq",
	}
	outputs := []string{
		"apbqcr",
		"apbqrs",
		"apbqcd",
	}
	for i := 0; i < len(inputs); i++ {
		result := mergeAlternately(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
