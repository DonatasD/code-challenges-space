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

func postorderTraversal(root *TreeNode) []int {
	var res []int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}
	dfs(root)
	return res
}

func main() {
	inputs := []*TreeNode{
		{
			Val: 1,
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 3},
			},
		},
		{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 4},
				Right: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			Right: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 9}},
			},
		},
		nil,
		{
			Val: 1,
		},
	}

	outputs := [][]int{
		{3, 2, 1},
		{4, 6, 7, 5, 2, 9, 8, 3, 1},
		{},
		{1},
	}
	for i := 0; i < len(inputs); i++ {
		result := postorderTraversal(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
