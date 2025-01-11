package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func countAndSay(n int) string {
	res := "1"
	var encrypt func(str string) string
	encrypt = func(str string) string {
		encrypted := ""
		count := 1
		cur := str[0:1]
		for i := 1; i < len(str); i++ {
			if cur == str[i:i+1] {
				count++
			} else {
				encrypted = encrypted + strconv.Itoa(count) + cur
				count = 1
				cur = str[i : i+1]
			}
		}
		return encrypted + strconv.Itoa(count) + cur
	}

	for i := 1; i < n; i++ {
		res = encrypt(res)
	}
	return res
}

func main() {
	inputs := []int{4, 1}
	outputs := []string{"1211", "1"}
	for i := 0; i < len(inputs); i++ {
		result := countAndSay(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
