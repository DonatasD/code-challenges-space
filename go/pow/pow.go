package main

import (
	"fmt"
	"reflect"
)

func myPow(x float64, n int) float64 {
	var calc func(x float64, n int) float64
	calc = func(x float64, n int) float64 {
		if x == 0 {
			return 0
		}
		if n == 0 {
			return 1
		}
		res := calc(x*x, n/2)
		if n%2 == 0 {
			return res
		}
		return x * res
	}
	res := calc(x, n)
	if n < 0 {
		return 1.0 / res
	}
	return res
}

func main() {
	inputs := []float64{2, 2.1, 2}
	params := []int{10, 3, -2}
	outputs := []float64{1024, 9.261, 0.25}
	for i := 0; i < len(inputs); i++ {

		result := myPow(inputs[i], params[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
