package main

import (
	"fmt"
	"reflect"
	"strings"
)

func solveNQueens(n int) [][]string {
	col := map[int]bool{}
	posDiag := map[int]bool{}
	negDiag := map[int]bool{}

	var res [][]string
	var board [][]string
	for i := 0; i < n; i++ {
		var row []string
		for j := 0; j < n; j++ {
			row = append(row, ".")
		}
		board = append(board, row)
	}
	var backtrack func(r int)
	backtrack = func(r int) {
		if r == n {
			var temp []string
			for i := 0; i < len(board); i++ {
				temp = append(temp, strings.Join(board[i], ""))
			}
			res = append(res, temp)
			return
		}

		for c := 0; c < n; c++ {
			colVal := col[c]
			posDiagVal := posDiag[r+c]
			negDiagVal := negDiag[r-c]
			if colVal || posDiagVal || negDiagVal {
				continue
			}
			col[c] = true
			posDiag[r+c] = true
			negDiag[r-c] = true
			board[r][c] = "Q"

			backtrack(r + 1)

			col[c] = false
			posDiag[r+c] = false
			negDiag[r-c] = false
			board[r][c] = "."
		}
	}
	backtrack(0)
	return res
}

func main() {
	inputs := []int{4, 1}
	outputs := [][][]string{
		{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		{{"Q"}},
	}

	for i := 0; i < len(inputs); i++ {

		result := solveNQueens(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
