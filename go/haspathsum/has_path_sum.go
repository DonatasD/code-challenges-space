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

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	newTargetSum := targetSum - root.Val
	if root.Right == nil && root.Left == nil {
		return newTargetSum == 0
	}
	return hasPathSum(root.Left, newTargetSum) || hasPathSum(root.Right, newTargetSum)
}

func main() {
	inputs := []*TreeNode{
		{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 11,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			Right: &TreeNode{
				Val:  8,
				Left: &TreeNode{Val: 13},
				Right: &TreeNode{
					Val:   4,
					Right: &TreeNode{Val: 1},
				},
			},
		},
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		nil,
	}

	params := []int{
		22,
		5,
		0,
	}

	outputs := []bool{
		true,
		false,
		false,
	}
	for i := 0; i < len(inputs); i++ {
		result := hasPathSum(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
