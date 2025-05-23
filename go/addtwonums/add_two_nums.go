package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var rem int = 0

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var val1, val2 int
	var next1, next2 *ListNode
	if l1 == nil {
		val1 = 0
		next1 = nil
	} else {
		val1 = l1.Val
		next1 = l1.Next
	}
	if l2 == nil {
		val2 = 0
		next2 = nil
	} else {
		val2 = l2.Val
		next2 = l2.Next
	}
	var sum = (val1 + val2 + rem) % 10
	rem = (val1 + val2 + rem) / 10
	if l1 == nil && l2 == nil {
		if sum == 0 {
			return nil
		} else {
			return &ListNode{sum, nil}
		}
	}
	return &ListNode{sum, addTwoNumbers(next1, next2)}
}

func main() {
	inputs := []*ListNode{{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	params := []*ListNode{{9, &ListNode{9, &ListNode{9, nil}}}}
	outputs := []*ListNode{{8, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}
	for i := 0; i < len(inputs); i++ {

		result := addTwoNumbers(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
