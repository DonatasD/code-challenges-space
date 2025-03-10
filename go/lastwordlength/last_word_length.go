package main

import (
	"fmt"
	"reflect"
)

func lengthOfLastWord(s string) int {
	trim := true
	res := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if trim {
				continue
			} else {
				break
			}
		}
		trim = false
		res++
	}
	return res
}

func main() {
	inputs := []string{
		"Hello World",
		"   fly me   to   the moon  ",
		"luffy is still joyboy",
	}
	outputs := []int{5, 4, 6}
	for i := 0; i < len(inputs); i++ {

		result := lengthOfLastWord(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}

	}
}
