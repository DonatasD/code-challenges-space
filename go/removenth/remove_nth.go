package main

import (
	"fmt"
	"reflect"
)

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	var arr []*ListNode
	cur := head
	arr = append(arr, cur)
	for cur.Next != nil {
		arr = append(arr, cur.Next)
		cur = cur.Next
	}
	if len(arr) == 1 {
		return nil
	} else if len(arr) == n {
		return arr[1]
	} else {
		index := len(arr) - n - 1
		el := arr[index]
		if el.Next.Next != nil {
			el.Next = el.Next.Next
		} else {
			el.Next = nil
		}
	}
	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	left, right := dummy, head

	for n > 0 && right != nil {
		right = right.Next
		n--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	inputs := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
		{Val: 1, Next: nil},
		{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
	}
	params := []int{
		2,
		1,
		1,
	}
	outputs := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}},
		nil,
		{Val: 1, Next: nil},
	}
	for i := 0; i < len(inputs); i++ {
		result := removeNthFromEnd(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
