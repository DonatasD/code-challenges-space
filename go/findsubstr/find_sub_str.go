package main

import (
	"fmt"
	"reflect"
)

func findSubstring(s string, words []string) []int {
	k := len(words)
	n := len(words[0])
	if len(s) < n {
		return []int{}
	}
	wordCounts := make(map[string]int, n)
	for _, w := range words {
		wordCounts[w]++
	}
	seenTimes := make(map[string]int, n)
	var head, qlen int
	queue := make([]string, 0, k)
	reset := func() {
		head, qlen = 0, 0
		queue = queue[:0]
		for k := range seenTimes {
			delete(seenTimes, k)
		}
	}
	res := []int{}
	for i := 0; i < n; i++ {
		reset()
		for j := i; j+n <= len(s); j += n {
			w := s[j : j+n]
			wordCount, exists := wordCounts[w]
			if !exists {
				reset()
				continue
			}
			for seenTimes[w] == wordCount {
				seenTimes[queue[head]]--
				head++
				qlen--
			}
			seenTimes[w]++
			queue = append(queue, w)
			qlen++
			if qlen == k {
				res = append(res, j-(k-1)*n)
			}
		}
	}
	return res
}

func main() {
	inputs := []string{
		"barfoothefoobarman",
		"wordgoodgoodgoodbestword",
		"barfoofoobarthefoobarman",
		"lingmindraboofooowingdingbarrwingmonkeypoundcake",
		"wordgoodgoodgoodbestword",
	}
	params := [][]string{
		{"foo", "bar"},
		{"word", "good", "best", "word"},
		{"bar", "foo", "the"},
		{"fooo", "barr", "wing", "ding", "wing"},
		{"word", "good", "best", "good"},
	}
	outputs := [][]int{
		{0, 9},
		{},
		{6, 9, 12},
		{13},
		{8},
	}

	for i := 0; i < len(inputs); i++ {
		result := findSubstring(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed: %v\n", i, result)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
