package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	res := &ListNode{Val: 0, Next: nil}
	var backtrack func(cur *ListNode, cur1 *ListNode, cur2 *ListNode)
	backtrack = func(cur *ListNode, cur1 *ListNode, cur2 *ListNode) {
		if cur1 == nil || cur2 != nil && cur2.Val <= cur1.Val {
			cur.Val = cur2.Val
			if cur2.Next == nil && cur1 == nil {
				return
			} else {
				cur.Next = &ListNode{Val: 0, Next: nil}
			}
			cur.Next = &ListNode{Val: 0, Next: nil}
			backtrack(cur.Next, cur1, cur2.Next)
		} else {
			cur.Val = cur1.Val
			if cur1.Next == nil && cur2 == nil {
				return
			} else {
				cur.Next = &ListNode{Val: 0, Next: nil}
			}
			backtrack(cur.Next, cur1.Next, cur2)
		}
	}
	backtrack(res, list1, list2)
	return res
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{Val: 0, Next: nil}
	tail := res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else if list2 != nil {
		tail.Next = list2
	}
	return res.Next
}

func main() {
	inputs := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}},
		nil,
		nil,
	}
	params := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
		nil,
		{Val: 0, Next: nil},
	}
	outputs := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}}},
		nil,
		{Val: 0, Next: nil},
	}
	for i := 0; i < len(inputs); i++ {
		result := mergeTwoLists(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
