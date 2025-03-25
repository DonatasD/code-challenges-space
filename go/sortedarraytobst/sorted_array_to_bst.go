package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	m := len(nums) / 2

	return &TreeNode{Val: nums[m], Left: sortedArrayToBST(nums[:m]), Right: sortedArrayToBST(nums[m+1:])}
}

func main() {
	inputs := [][]int{
		{-10, -3, 0, 5, 9},
		{1, 3},
	}
	outputs := []*TreeNode{
		{Val: 0, Left: &TreeNode{Val: -3, Left: &TreeNode{Val: -10}}, Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}}},
		{Val: 3, Left: &TreeNode{Val: 1}},
	}
	for i := 0; i < len(inputs); i++ {
		result := sortedArrayToBST(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
