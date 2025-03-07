package main

import (
	"fmt"
	"reflect"
)

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return NumArray{prefixSum: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.prefixSum[right]
	}

	return this.prefixSum[right] - this.prefixSum[left-1]
}

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	inputs := [][]int{{0, 2}, {2, 5}, {0, 5}}
	outputs := []int{1, -1, -3}
	for i := 0; i < len(inputs); i++ {

		result := obj.SumRange(inputs[i][0], inputs[i][1])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}

	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
