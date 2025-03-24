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

func inorderTraversal(root *TreeNode) []int {
	var res []int

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			dfs(root.Left)
		}
		res = append(res, root.Val)
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	if root != nil {
		dfs(root)
	}
	return res
}

func main() {
	inputs := []*TreeNode{
		{
			Val: 1,
			Right: &TreeNode{
				Left: &TreeNode{
					Val: 3,
				},
				Val: 2,
			},
		},
		{
			Left: &TreeNode{
				Left: &TreeNode{
					Val: 4,
				},
				Val: 2,
				Right: &TreeNode{
					Left: &TreeNode{
						Val: 6,
					},
					Val: 5,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			Val: 1,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Left: &TreeNode{
						Val: 9,
					},
					Val: 8,
				},
			},
		},
		nil,
		{Val: 1},
	}

	outputs := [][]int{
		{1, 3, 2},
		{4, 2, 6, 5, 7, 1, 3, 9, 8},
		{},
		{1},
	}
	for i := 0; i < len(inputs); i++ {
		result := inorderTraversal(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
