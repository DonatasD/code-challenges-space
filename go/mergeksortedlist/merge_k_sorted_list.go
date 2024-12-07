package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists1(lists []*ListNode) *ListNode {
	res := &ListNode{Val: 0, Next: nil}
	tail := res
	for {
		first := true
		found := false
		maxIndex := 0
		for i, v := range lists {
			if v != nil && (first || lists[i].Val < lists[maxIndex].Val) {
				first = false
				maxIndex = i
				found = true
			}
		}
		if !found {
			break
		} else {
			tail.Next = lists[maxIndex]
			lists[maxIndex] = lists[maxIndex].Next
			tail = tail.Next
		}
	}

	return res.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	var mergeList func(l1 *ListNode, l2 *ListNode) *ListNode
	mergeList = func(l1 *ListNode, l2 *ListNode) *ListNode {
		dummy := &ListNode{Val: 0, Next: nil}
		tail := dummy

		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				tail.Next = l1
				l1 = l1.Next
			} else {
				tail.Next = l2
				l2 = l2.Next
			}
			tail = tail.Next
		}

		if l1 != nil {
			tail.Next = l1
		} else if l2 != nil {
			tail.Next = l2
		}
		return dummy.Next
	}

	if lists == nil || len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		var mergedList []*ListNode

		for i := 0; i < len(lists); i += 2 {
			l1 := lists[i]
			var l2 *ListNode = nil
			if i+1 < len(lists) {
				l2 = lists[i+1]
			}
			mergedList = append(mergedList, mergeList(l1, l2))
		}
		lists = mergedList
	}
	return lists[0]
}

func main() {
	inputs := [][]*ListNode{
		{
			{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}},
			{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
			{Val: 2, Next: &ListNode{Val: 6, Next: nil}},
		},
		{},
		{nil},
	}
	outputs := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}}},
		nil,
		nil,
	}
	for i := 0; i < len(inputs); i++ {
		result := mergeKLists(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
