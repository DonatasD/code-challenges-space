package main

import (
	"fmt"
	"reflect"
)

func maxProfit(prices []int) int {
	minPrice, res, cur := 100000, 0, 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		cur = prices[i] - minPrice
		if res < cur {
			res = cur
		}
	}
	return res
}

func main() {
	inputs := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
		{1, 2},
	}
	outputs := []int{
		5, 0, 1,
	}
	for i := 0; i < len(inputs); i++ {

		result := maxProfit(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
