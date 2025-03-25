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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Right), maxDepth(root.Left))
}

func main() {
	inputs := []*TreeNode{
		{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		},
		{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
	}

	outputs := []int{
		3,
		2,
	}
	for i := 0; i < len(inputs); i++ {
		result := maxDepth(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
