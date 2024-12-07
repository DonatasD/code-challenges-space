package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	if head == nil {
		return nil
	}
	var backtrack func(res *ListNode, h *ListNode)
	backtrack = func(res *ListNode, h *ListNode) {
		if h.Next == nil {
			res.Next = &ListNode{Val: h.Val, Next: nil}
		} else {
			next := &ListNode{Val: h.Val, Next: nil}
			res.Next = &ListNode{Val: h.Next.Val, Next: next}
			if h.Next.Next != nil {
				backtrack(res.Next.Next, h.Next.Next)
			}
		}
	}
	backtrack(dummy, head)
	return dummy.Next
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prev, curr := dummy, head

	for curr != nil && curr.Next != nil {
		nextPair := curr.Next.Next
		second := curr.Next

		second.Next = curr
		curr.Next = nextPair
		prev.Next = second

		prev = curr
		curr = nextPair
	}
	return dummy.Next
}

func main() {
	inputs := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}},
		nil,
		{Val: 1, Next: nil},
		{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}},
	}
	outputs := []*ListNode{
		{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}},
		nil,
		{Val: 1, Next: nil},
		{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}},
	}
	for i := 0; i < len(inputs); i++ {
		result := swapPairs(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
