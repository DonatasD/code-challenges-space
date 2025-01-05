package main

import (
	"fmt"
	"reflect"
)

func isValidSudoku(board [][]byte) bool {
	cols := make(map[int]map[byte]bool)
	rows := make(map[int]map[byte]bool)
	squares := make(map[string]map[byte]bool)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			val, sqKey := board[r][c], fmt.Sprintf("%d-%d", r/3, c/3)
			if val == '.' {
				continue
			}
			_, cExists := cols[c]
			_, rExists := rows[r]
			_, sExists := squares[sqKey]
			if !cExists {
				cols[c] = make(map[byte]bool)
			}
			if !rExists {
				rows[r] = make(map[byte]bool)
			}
			if !sExists {
				squares[sqKey] = make(map[byte]bool)
			}
			_, cmExists := cols[c][val]
			_, rmExists := rows[r][val]
			_, smExists := squares[sqKey][val]
			if cmExists || rmExists || smExists {
				return false
			}
			cols[c][val] = true
			rows[r][val] = true
			squares[sqKey][val] = true
		}
	}
	return true
}

func main() {
	inputs := [][][]byte{
		{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
		{
			{'8', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
	}
	outputs := []bool{true, false}
	for i := 0; i < len(inputs); i++ {
		result := isValidSudoku(inputs[i])
		if reflect.DeepEqual(result, outputs[i]) {
			fmt.Printf("Testcase %v passed\n", i)
		} else {
			fmt.Printf("failed %v: %v expected: %v\n", i, result, outputs[i])
		}
	}
}
