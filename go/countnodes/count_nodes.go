package main

import (
	"fmt"
	"math"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	var dfsL, dfsR func(root *TreeNode) int
	dfsL = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return 1 + dfsL(root.Left)
	}
	dfsR = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return 1 + dfsR(root.Right)
	}
	l, r := dfsL(root), dfsR(root)
	if l > r {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	}
	return int(math.Pow(float64(2), float64(l))) - 1
}

func main() {
	inputs := []*TreeNode{
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 6,
				},
			},
		},
		nil,
		{
			Val: 1,
		},
	}
	outputs := []int{6, 0, 1}
	for i := 0; i < len(inputs); i++ {

		result := countNodes(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
