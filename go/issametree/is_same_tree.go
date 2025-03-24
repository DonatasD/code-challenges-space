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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	inputs := [][]*TreeNode{
		{
			{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		},
		{
			{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
			{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
		},
	}

	outputs := []bool{
		true,
		false,
		false,
	}
	for i := 0; i < len(inputs); i++ {
		result := isSameTree(inputs[i][0], inputs[i][1])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
