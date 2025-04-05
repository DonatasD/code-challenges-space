package main

import (
	"fmt"
	"reflect"
)

func totalNQueens(n int) int {
	var count int
	var solve func(row, columns, diagonals1, diagonals2 int)

	solve = func(row, columns, diagonals1, diagonals2 int) {
		if row == n {
			count++
			return
		}

		availablePositions := ((1 << n) - 1) & ^(columns | diagonals1 | diagonals2)
		for availablePositions != 0 {
			position := availablePositions & -availablePositions

			availablePositions ^= position
			solve(row+1, columns|position,
				(diagonals1|position)<<1,
				(diagonals2|position)>>1)
		}
	}

	solve(0, 0, 0, 0)
	return count
}

func main() {
	inputs := []int{4, 1}
	outputs := []int{2, 1}

	for i := 0; i < len(inputs); i++ {

		result := totalNQueens(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
