package main

import (
	"fmt"
	"reflect"
)

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0, len(strs))
	countMap := make(map[[26]byte][]string, len(strs))

	for _, str := range strs {
		count := [26]byte{}
		for _, s := range str {
			count[s-'a']++
		}
		countMap[count] = append(countMap[count], str)
	}
	for _, strings := range countMap {
		result = append(result, strings)
	}

	return result
}

func main() {
	inputs := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""},
		{"a"},
		{"bdddddddddd", "bbbbbbbbbbc"},
	}
	outputs := [][][]string{
		{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		{{""}},
		{{"a"}},
		{{"bdddddddddd"}, {"bbbbbbbbbbc"}},
	}
	for i := 0; i < len(inputs); i++ {

		result := groupAnagrams(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
