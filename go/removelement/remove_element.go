package main

import "fmt"

func removeElement(nums []int, val int) int {
	l := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != val {
			nums[l] = nums[r]
			l++
		}
	}
	return l
}

func main() {
	inputs := [][]int{{3, 2, 2, 3}, {0, 1, 2, 2, 3, 0, 4, 2}}
	params := []int{3, 2}
	outputs := []int{2, 5}

	for i := range inputs {

		result := removeElement(inputs[i], params[i])
		if result == outputs[i] {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
