package main

import "fmt"

func maxArea(height []int) int {
	maxArea := 0

	for i, j := 0, len(height)-1; i != j; {
		l, r, length, minValue := height[i], height[j], j-i, 0
		if l >= r {
			minValue = r
			j--
		} else {
			minValue = l
			i++
		}
		curAreaMax := minValue * length
		if curAreaMax > maxArea {
			maxArea = curAreaMax
		}
	}

	return maxArea
}

func main() {
	inputs := [][]int{{1, 8, 6, 2, 5, 4, 8, 3, 7}, {1, 1}}
	outputs := []int{49, 1}

	for i, v := range inputs {

		result := maxArea(v)
		if result == outputs[i] {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
