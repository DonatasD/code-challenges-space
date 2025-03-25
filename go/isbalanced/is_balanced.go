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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isBalanced(root *TreeNode) bool {
	var dfs func(root *TreeNode) (bool, int)
	dfs = func(root *TreeNode) (bool, int) {
		if root == nil {
			return true, 0
		}
		isBalancedLeft, leftHeight := dfs(root.Left)
		isBalancedRight, rightHeight := dfs(root.Right)
		return isBalancedLeft && isBalancedRight && abs(leftHeight-rightHeight) <= 1, 1 + max(leftHeight, rightHeight)
	}
	res, _ := dfs(root)
	return res
}

func main() {
	inputs := []*TreeNode{
		{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 8}},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}},
		},
		{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
		},
		{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 2},
		},
		nil,
	}

	outputs := []bool{
		true,
		true,
		false,
		true,
	}
	for i := 0; i < len(inputs); i++ {
		result := isBalanced(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
