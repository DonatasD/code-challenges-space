package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var dfs func(node *ListNode) bool
	dfs = func(node *ListNode) bool {
		if node == nil {
			return false
		}
		if node.Next == head {
			return true
		}
		next := node.Next
		node.Next = head
		return dfs(next)
	}
	return dfs(head)
}

func main() {
	input1 := &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4, Next: nil}}}
	input1.Next.Next.Next = input1
	input2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	input2.Next.Next = input2
	inputs := []*ListNode{
		{Val: 3, Next: input1},
		input2,
		{Val: 1, Next: nil},
	}

	outputs := []bool{true, true, false}
	for i := 0; i < len(inputs); i++ {
		result := hasCycle(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
