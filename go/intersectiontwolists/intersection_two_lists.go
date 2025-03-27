package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	node1, node2 := headA, headB
	for node1 != node2 {
		if node1 != nil {
			node1 = node1.Next
		} else {
			node1 = headB
		}
		if node2 != nil {
			node2 = node2.Next
		} else {
			node2 = headA
		}
	}
	return node1
}

func main() {
	common1 := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	common2 := &ListNode{Val: 2, Next: &ListNode{Val: 4}}
	inputs := [][]*ListNode{
		{
			{Val: 4, Next: &ListNode{Val: 1, Next: common1}},
			{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: common1}}},
		},
		{
			{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: common2}}},
			{Val: 3, Next: common2},
		},
		{
			{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			{Val: 1, Next: &ListNode{Val: 5}},
		},
	}
	outputs := []*ListNode{
		common1,
		common2,
		nil,
	}
	for i := 0; i < len(inputs); i++ {
		result := getIntersectionNode(inputs[i][0], inputs[i][1])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
