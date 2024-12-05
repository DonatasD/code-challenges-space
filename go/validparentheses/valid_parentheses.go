package main

import (
	"fmt"
	"reflect"
)

func isValid(s string) bool {
	res := ""
	c := map[string]string{"}": "{", ")": "(", "]": "["}
	for i := 0; i < len(s); i++ {
		cur := s[i : i+1]
		switch cur {
		case "}", ")", "]":
			open := c[cur]
			if len(res) == 0 {
				return false
			}
			last := res[len(res)-1:]
			if last != open {
				return false
			}
			res = res[0 : len(res)-1]
		default:
			res += cur
		}
	}

	return len(res) == 0
}

func main() {
	inputs := []string{"()", "()[]{}", "(]", "([])"}
	outputs := []bool{true, true, false, true}
	for i := 0; i < len(inputs); i++ {
		result := isValid(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
